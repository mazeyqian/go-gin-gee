GOOS=darwin GOARCH=amd64 go build -o dist/batch-git-pull-mac-darwin-amd64 scripts/batch-git-pull/main.go

GOOS=linux GOARCH=amd64 go build -o dist/batch-git-pull-linux-amd64 scripts/batch-git-pull/main.go

# https://linuxhint.com/check_memory_usage_process_linux/
ps -o pid,user,%mem,command ax | sort -b -k3 -r
