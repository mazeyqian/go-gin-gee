# go-gin-gee

## Install

```
go get
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn
```

## Run

```
go run cmd/api/main.go
```

Visit: `http://127.0.0.1:3000/api/ping`

## Build

```
go build main.go
```
