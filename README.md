# go-gin-gee

## Install

```
go get

# or
go get github.com/bitfield/script
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn
```

## Run

```
go run scripts/init/main.go

go run cmd/api/main.go

# or
go run scripts/change-git-user/main.go
```

Visit: `http://127.0.0.1:3000/api/ping`

## Build

### Default

```
go build cmd/api/main.go
```

### Linux

```
GOOS=linux GOARCH=amd64 go build -o dist/api cmd/api/main.go

# or
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go
```

### Mac

```
# Api
GOOS=darwin GOARCH=amd64 go build -o dist/api-mac-darwin-amd64 cmd/api/main.go

# Init
GOOS=darwin GOARCH=amd64 go build -o dist/init-mac-darwin-amd64 scripts/init/main.go

# Change git user
GOOS=darwin GOARCH=amd64 go build -o dist/change-git-user-mac-darwin-amd64 scripts/change-git-user/main.go
```

### Windows

```
GOOS=windows GOARCH=amd64 go build -o dist/api-windows-amd64 cmd/api/main.go
```

### Config

```
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api
autostart=true
autorestart=true
stderr_logfile=/web/go-gin-gee/log/api.err
stdout_logfile=/web/go-gin-gee/log/api.log
```

### Scripts

Init

```
go run scripts/init/main.go
```

Git Pull

```
go run scripts/batch-git-pull/main.go
```
