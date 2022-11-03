# go-gin-gee

Gee provides several services for everyday life. The project is based on [gin](https://github.com/gin-gonic/gin), and the structure refers to [project-layout](https://github.com/golang-standards/project-layout). There are some daily scripts in the folder `scripts` depending on [script](https://github.com/bitfield/script), which can run by the command `go run`.

**Table of Contents**

- [go-gin-gee](#go-gin-gee)
  - [Install](#install)
  - [Run](#run)
  - [Build](#build)
    - [Default](#default)
    - [Linux](#linux)
    - [Mac](#mac)
    - [Windows](#windows)
  - [Supervisor Config](#supervisor-config)
  - [Scripts](#scripts)

## Install

```
# All Dependences
go mod download

# Add/Update
go install github.com/example/name
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn
```

## Run

It's necessary to run the command `go run scripts/init/main.go` when serving the project first.

```
# Serve
go run cmd/api/main.go

# Restart
# cd /web/go-gin-gee
go run scripts/restart/main.go
```

Visit: `http://127.0.0.1:3000/api/ping`.

```
pong/v1.0.0/2022-09-29 04:52:43
```

## Build

### Default

```
go build cmd/api/main.go
```

### Linux

It's usually useful to run the command `chmod u+x script-name-linux-amd64` if the permission error happens.

```
# API
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go

# Scripts
GOOS=linux GOARCH=amd64 go build -o dist/init-linux-amd64 scripts/init/main.go
```

### Mac

```
# API
GOOS=darwin GOARCH=amd64 go build -o dist/api-mac-darwin-amd64 cmd/api/main.go

# Scripts
GOOS=darwin GOARCH=amd64 go build -o dist/init-mac-darwin-amd64 scripts/init/main.go
```

### Windows

```
# API
GOOS=windows GOARCH=amd64 go build -o dist/api-windows-amd64 cmd/api/main.go
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
go run scripts/change-git-user/main.go
```

Git pull all projects in a folder.

```
go run scripts/batch-git-pull/main.go
```

Transfer apple note table to markdown table. 

```
go run scripts/transfer-notes-to-md-table/main.go
```

More in folder `scripts`.
