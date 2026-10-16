# pwsh -NoProfile -File scripts/powershell/test/test-batch-remove-orphan-lrc.ps1
$ErrorActionPreference = 'Stop'
$scriptPath = Join-Path (Split-Path -Parent $PSScriptRoot) 'batch-remove-orphan-lrc.ps1'
$engine = (Get-Process -Id $PID).Path
$root = Join-Path ([System.IO.Path]::GetTempPath()) ('orphan-lrc-test-' + [guid]::NewGuid().ToString('N'))

function New-File {
  param([string]$Directory, [string]$Name)
  [System.IO.File]::WriteAllText((Join-Path $Directory $Name), $Name)
}

function Assert-File {
  param([string]$Directory, [string]$Name, [bool]$Exists = $true)
  $file = Join-Path $Directory $Name
  if ([System.IO.File]::Exists($file) -ne $Exists) { throw "Unexpected existence: $file" }
  if ($Exists -and [System.IO.File]::ReadAllText($file) -cne $Name) { throw "Content changed: $file" }
}

function Invoke-Cleanup {
  param([int]$Status, [string[]]$Arguments)
  $ErrorActionPreference = 'Continue'
  $output = & $engine -NoProfile -ExecutionPolicy Bypass -File $scriptPath @Arguments 2>&1
  $ErrorActionPreference = 'Stop'
  if ($LASTEXITCODE -ne $Status) { throw "Expected $Status, got ${LASTEXITCODE}: $output" }
  if ($Status -ne 0) {
    if (!$output) { throw 'Failure did not report an error' }
    if (($output | Out-String) -match 'Task Complete!|Preview Complete!') {
      throw "Failure printed a success summary: $output"
    }
  }
}

try {
  $music = Join-Path $root 'music [test]'
  $child = Join-Path $music 'child'
  $null = New-Item -ItemType Directory -Path $child -Force
  $extensions = @('mp3', 'ogg', 'oga', 'opus', 'flac', 'wav', 'wave', 'm4a', 'aac', 'wma', 'aif', 'aiff', 'aifc', 'ape', 'alac', 'wv', 'dsf', 'dff')
  foreach ($ext in $extensions) {
    New-File $music ("Track-$ext." + $ext.ToUpperInvariant())
    New-File $music "track-$ext.LRC"
  }
  $unicodeName = ([char]0x6B4C).ToString() + ' [live].part'
  foreach ($name in @('orphan.lrc', 'orphan [1].LRC', 'parent-only.mp3', 'child-only.lrc', 'Song.mp3', 'Song-live.lrc', 'Song .lrc', 'notes.txt', 'unsupported.xyz', 'unsupported.lrc', '.hidden.mp3', '.hidden.lrc', '.orphan.lrc', "$unicodeName.flac", "$unicodeName.lrc")) {
    New-File $music $name
  }
  foreach ($name in @('child-only.ogg', 'parent-only.lrc', 'nested-orphan.lrc', 'match.mp3', 'match.lrc')) { New-File $child $name }
  $null = New-Item -ItemType Directory -Path (Join-Path $music 'directory.lrc')

  # On Windows, exercise the hidden attribute as well as dot-prefixed names.
  if ($env:OS -eq 'Windows_NT') {
    $hidden = Get-Item -LiteralPath (Join-Path $music '.hidden.mp3') -Force
    $hidden.Attributes = $hidden.Attributes -bor [System.IO.FileAttributes]::Hidden
  }

  Invoke-Cleanup 0 @('-Path', $music, '-Recurse', '-WhatIf')
  Assert-File $music 'orphan.lrc'
  Assert-File $child 'nested-orphan.lrc'
  Invoke-Cleanup 0 @('-Path', $music)
  foreach ($name in @('orphan.lrc', 'orphan [1].LRC', 'child-only.lrc', 'Song-live.lrc', 'Song .lrc', 'unsupported.lrc', '.orphan.lrc')) { Assert-File $music $name $false }
  Assert-File $child 'nested-orphan.lrc'
  Invoke-Cleanup 0 @('-Path', $music, '-Recurse')
  Assert-File $child 'nested-orphan.lrc' $false
  Assert-File $child 'parent-only.lrc' $false
  foreach ($ext in $extensions) {
    Assert-File $music ("Track-$ext." + $ext.ToUpperInvariant())
    Assert-File $music "track-$ext.LRC"
  }
  foreach ($name in @('parent-only.mp3', 'Song.mp3', 'notes.txt', 'unsupported.xyz', '.hidden.mp3', '.hidden.lrc', "$unicodeName.flac", "$unicodeName.lrc")) { Assert-File $music $name }
  foreach ($name in @('child-only.ogg', 'match.mp3', 'match.lrc')) { Assert-File $child $name }
  if (!(Test-Path -LiteralPath (Join-Path $music 'directory.lrc') -PathType Container)) { throw 'Directory removed' }
  Invoke-Cleanup 0 @('-Path', $music, '-Recurse')
  $empty = Join-Path $root 'empty'
  $null = New-Item -ItemType Directory -Path $empty
  Invoke-Cleanup 0 @('-Path', $empty)
  Invoke-Cleanup 1 @('-Path', (Join-Path $root 'missing'))
  Invoke-Cleanup 1 @('-Path', (Join-Path $music 'Song.mp3'))

  if ($env:OS -ne 'Windows_NT') {
    $outside = Join-Path $root 'outside'
    $null = New-Item -ItemType Directory -Path $outside
    New-File $outside 'linked.lrc'
    $directoryLink = Join-Path $music 'linked-directory'
    $fileLink = Join-Path $music 'linked.lrc'
    $null = New-Item -ItemType SymbolicLink -Path $directoryLink -Target $outside
    $null = New-Item -ItemType SymbolicLink -Path $fileLink -Target (Join-Path $outside 'linked.lrc')
    Invoke-Cleanup 0 @('-Path', $music, '-Recurse')
    Assert-File $outside 'linked.lrc'
    if (!(Test-Path -LiteralPath $fileLink)) { throw 'LRC link removed' }
    Invoke-Cleanup 1 @('-Path', $directoryLink, '-Recurse')

    if ((& id -u) -ne '0') {
      $failureRoot = Join-Path $root 'scan-failure'
      $blocked = Join-Path $failureRoot 'blocked'
      $null = New-Item -ItemType Directory -Path $blocked -Force
      New-File $failureRoot 'must-stay.lrc'
      try {
        & chmod 000 $blocked
        if ($LASTEXITCODE -ne 0) { throw 'Could not set scan-failure permissions' }
        Invoke-Cleanup 1 @('-Path', $failureRoot, '-Recurse')
        Assert-File $failureRoot 'must-stay.lrc'
      } finally {
        & chmod 700 $blocked
        if ($LASTEXITCODE -ne 0) { throw 'Could not restore directory permissions' }
      }

      $deleteFailure = Join-Path $root 'delete-failure'
      $null = New-Item -ItemType Directory -Path $deleteFailure
      New-File $deleteFailure 'must-stay.lrc'
      try {
        & chmod 500 $deleteFailure
        if ($LASTEXITCODE -ne 0) { throw 'Could not set deletion-failure permissions' }
        Invoke-Cleanup 1 @('-Path', $deleteFailure)
        Assert-File $deleteFailure 'must-stay.lrc'
      } finally {
        & chmod 700 $deleteFailure
        if ($LASTEXITCODE -ne 0) { throw 'Could not restore directory permissions' }
      }
    } else {
      Write-Host 'Skipping permission-denial checks when running as root.'
    }
  }
  Write-Host 'All orphan LRC regression checks passed.'
} finally {
  if (Test-Path -LiteralPath $root) { Remove-Item -LiteralPath $root -Recurse -Force }
}
