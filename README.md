# go-gin-gee

Gee provides several services of everyday life. The structure refers to [project-layout](https://github.com/golang-standards/project-layout).

## Install

```
# All Dependences
go mod download;

# Single
go get github.com/example/name;

# Update
go get -u github.com/example/name;
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn;
```

## Run

```
# Init
go run scripts/init/main.go;

# Serve
go run cmd/api/main.go;

# Restart
# cd /web/go-gin-gee;
go run scripts/restart/main.go;
```

Visit: `http://127.0.0.1:3000/api/ping`.

```
pong/2022-09-29 04:52:43
```

## Build

### Default

```
go build cmd/api/main.go;
```

### Linux

```
# Default
GOOS=linux GOARCH=amd64 go build -o dist/api cmd/api/main.go;

# Rename Output
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go;
```

### Mac

```
# Api
GOOS=darwin GOARCH=amd64 go build -o dist/api-mac-darwin-amd64 cmd/api/main.go;

# Init
GOOS=darwin GOARCH=amd64 go build -o dist/init-mac-darwin-amd64 scripts/init/main.go;

# Change Git User
GOOS=darwin GOARCH=amd64 go build -o dist/change-git-user-mac-darwin-amd64 scripts/change-git-user/main.go;
```

### Windows

```
GOOS=windows GOARCH=amd64 go build -o dist/api-windows-amd64 cmd/api/main.go;
```

## Supervisor Config

```
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api-linux-amd64
autostart=true
autorestart=true
stderr_logfile=/web/go-gin-gee/log/api.err
stdout_logfile=/web/go-gin-gee/log/api.log
```

## Scripts


Change git user and e-mail in a folder.

```
go run scripts/change-git-user/main.go;
```

Git pull all projects in a folder.

```
go run scripts/batch-git-pull/main.go;
```

Transfer apple note table to markdown table. 

```
go run scripts/transfer-notes-to-md-table/main.go;
```
