# Permanently delete LRC files without same-directory audio matches.
# pwsh -NoProfile -File scripts/powershell/batch-remove-orphan-lrc.ps1 -Path ./music -Recurse -WhatIf
# powershell.exe -NoProfile -File scripts\powershell\batch-remove-orphan-lrc.ps1 -Path E:\Music -WhatIf

[CmdletBinding(SupportsShouldProcess = $true)]
param(
  [Parameter(Mandatory = $true)]
  [string]$Path,

  [switch]$Recurse
)

$ErrorActionPreference = 'Stop'
$audioExtensions = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::OrdinalIgnoreCase)
foreach ($extension in @('mp3', 'ogg', 'oga', 'opus', 'flac', 'wav', 'wave', 'm4a', 'aac', 'wma', 'aif', 'aiff', 'aifc', 'ape', 'alac', 'wv', 'dsf', 'dff')) {
  $null = $audioExtensions.Add('.' + $extension)
}

$root = Get-Item -LiteralPath $Path -Force
if ($root -isnot [System.IO.DirectoryInfo]) {
  throw "Path is not a filesystem directory: $Path"
}
if ($root.Attributes -band [System.IO.FileAttributes]::ReparsePoint) {
  throw "Use a real directory, not a directory link: $Path"
}

$directories = [System.Collections.Generic.Queue[string]]::new()
$directories.Enqueue($root.FullName)
$orphans = [System.Collections.Generic.List[string]]::new()

# Finish every directory scan before deleting anything. Do not traverse directory links.
while ($directories.Count -gt 0) {
  $entries = @(Get-ChildItem -LiteralPath $directories.Dequeue() -Force)
  $audioNames = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::OrdinalIgnoreCase)
  foreach ($entry in $entries) {
    if ($entry.PSIsContainer) {
      if ($Recurse -and !($entry.Attributes -band [System.IO.FileAttributes]::ReparsePoint)) {
        $directories.Enqueue($entry.FullName)
      }
    } elseif ($audioExtensions.Contains($entry.Extension)) {
      $null = $audioNames.Add($entry.BaseName)
    }
  }
  foreach ($entry in $entries) {
    if (!$entry.PSIsContainer -and $entry.Extension -ieq '.lrc' -and
        !($entry.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -and
        !$audioNames.Contains($entry.BaseName)) {
      $orphans.Add($entry.FullName)
    }
  }
}

$deleted = 0
foreach ($orphan in $orphans) {
  if ($PSCmdlet.ShouldProcess($orphan, 'Permanently delete orphan LRC file')) {
    # File.Delete cannot recursively remove a directory if the path changes after scanning.
    [System.IO.File]::Delete($orphan)
    $deleted++
    Write-Host ("Deleted: {0}" -f $orphan)
  }
}
if ($WhatIfPreference) {
  Write-Host ("Preview Complete! Would delete {0} orphan LRC file(s)." -f $orphans.Count)
} else {
  Write-Host ("Task Complete! Deleted {0} orphan LRC file(s)." -f $deleted)
}
