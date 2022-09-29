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
```

### Mac

```
GOOS=darwin GOARCH=amd64 go build cmd/api/main.go

# or
GOOS=darwin GOARCH=amd64 go build -o dist/change-git-user-mac scripts/change-git-user/main.go

GOOS=darwin GOARCH=amd64 go build -o dist/init scripts/init/main.go
```

### Windows

```
GOOS=windows GOARCH=amd64 go build cmd/api/main.go
```

### Deploy

```
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api
autostart=true
autorestart=true
stderr_logfile=/web/go-gin-gee/log/api.err
stdout_logfile=/web/go-gin-gee/log/api.log
```

### scripts

Git pull.

```
go run scripts/batch-git-pull/main.go
```
