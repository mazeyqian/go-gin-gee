# go-gin-gee

Gee provides several services for everyday work. The project is based on Gin [1], and the structure refers to ProjectLayout [3]. In addition, some daily scripts in the folder Scripts depend on Script [4], which can be used by the command `run`.

**Table of Contents**

- [go-gin-gee](#go-gin-gee)
  - [Install](#install)
  - [Script Examples](#script-examples)
  - [API Examples](#api-examples)
    - [Generate Short Link](#generate-short-link)
    - [Save Data](#save-data)
    - [Get Data](#get-data)
  - [Build](#build)
    - [Linux](#linux)
    - [Mac](#mac)
    - [Windows](#windows)
  - [Deploy](#deploy)
    - [Supervisor](#supervisor)
    - [Docker](#docker)
      - [Build](#build-1)
      - [Run](#run)
  - [Contributing](#contributing)
  - [References](#references)

## Install

```
git clone https://github.com/mazeyqian/go-gin-gee.git
```

## Script Examples

1\. Change Git name and email for different projects.

```
go run scripts/change-git-user/main.go -path="/Users/X/Web" -username="Your Name" -useremail="your@email.com"
```

Usage: [English](https://github.com/mazeyqian/go-gin-gee/releases/tag/v1.0.0) | [简体中文](https://blog.mazey.net/2956.html)

2\. `git pull` all projects in a folder.

```
go run scripts/batch-git-pull/main.go -path="/Users/X/Web"
```

Usage: [English](https://github.com/mazeyqian/go-gin-gee/releases/tag/v1.1.0) | [简体中文](https://blog.mazey.net/3035.html)

3\. Transfer apple note table to markdown table. 

```
go run scripts/transfer-notes-to-md-table/main.go
```

4\. Convert Markdown text to TypeDoc comments.

```
go run scripts/convert-markdown-to-comments/main.go
```

More in folder `scripts`.

## API Examples

Domain is `https://feperf.com`.

### Generate Short Link

**Description:**

Generate the short link for the original link.

**Path:** **/api/gee/generate-short-link**

**Method:** **POST**

**Params:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| ori_link | string | Original Link | Yes |

**Example:**

```shell
curl --location --request POST 'https://feperf.com/api/gee/generate-short-link' \
--header 'Content-Type: application/json' \
--data-raw '{
	"ori_link": "https://blog.mazey.net/tiny?ts=654321-221467-f22c24-493220-228e97-d90c73"
}'
```

**Returns:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| tiny_link | string | Short Link | Yes |

**Example:**

Success: Status Code 201

```json
{
    "tiny_link": "https://feperf.com/t/b"
}
```

Failure: Status Code 400

```json
{
    "code": 400
}
```

### Save Data

**Description:**

Save the data for searching.

**Path:** **/api/gee/create-alias2data**
 
**Method:** **POST**

**Params:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| alias | string | Alias | Yes |
| data | string | Data | Yes |
| public | bool | Public | Yes |

**Example:**

```shell
curl --location --request POST 'https://feperf.com/api/gee/create-alias2data' \
--header 'Content-Type: application/json' \
--data-raw '{
    "alias": "alias example",
    "data": "data example",
    "public": true
}'
```

**Returns:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| id | int | ID | Yes |
| alias | string | Alias | Yes |
| data | string | Data | Yes |

**Example:**

Success: Status Code 201

```json
{
    "id": 2,
    "created_at": "2023-01-07T11:14:24.572495702+08:00",
    "updated_at": "2023-01-07T11:14:24.57882362+08:00",
    "alias": "alias example",
    "data": "data example"
}
```

Failure: Status Code 400

```json
{
    "code": 400,
    "message": "data exist"
}
```

### Get Data

**Description:**

Get the data.

**Path:** **/api/gee/get-data-by-alias**
 
**Method:** **GET**

**Params:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| alias | string | Alias | Yes |

**Example:**

```shell
curl --location 'https://feperf.com/api/gee/get-data-by-alias?alias=alias%20example'
```

**Returns:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| id | int | ID | Yes |
| alias | string | Alias | Yes |
| data | string | Data | Yes |

**Example:**

Success: Status Code 200

```json
{
    "data": {
        "id": 5,
        "created_at": "2023-05-16T13:46:10.518769+08:00",
        "updated_at": "2023-05-16T13:46:10.520977+08:00",
        "alias": "alias example",
        "data": "data example",
        "public": true
    }
}
```

Failure: Status Code 404

```json
{
    "code": 404,
    "message": "data not found"
}
```

## Build

Default:

```
go build cmd/api/main.go
```

### Linux

It's usually useful to run the command `chmod u+x script-name-linux-amd64` if the permission error happens.

```
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go
```

### Mac

```
GOOS=darwin GOARCH=amd64 go build -o dist/api-mac-darwin-amd64 cmd/api/main.go
```

### Windows

```
GOOS=windows GOARCH=amd64 go build -o dist/api-windows-amd64 cmd/api/main.go
```

## Deploy

### Supervisor

```
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api-linux-amd64 -configpath="data/config.json"
autostart=true
autorestart=true
```

### Docker

#### Build

```
# Command
bash ./scripts/docker-build.sh "{RUN_FLAG}" "WECOM_ROBOT_CHECK={WECOM_ROBOT_CHECK}"

# Example 1: Build
# RUN_FLAG is optional, default is "ONLY_BUILD"
# WECOM_ROBOT_CHECK is unnecessary.
bash ./scripts/docker-build.sh "ONLY_BUILD"

# Example 2: Build and Run
# RUN_FLAG is required.
# WECOM_ROBOT_CHECK is optional.
bash ./scripts/docker-build.sh "RUN" "WECOM_ROBOT_CHECK=b2d57746-7146-44f2-8207-86cb0ca832be"
```

#### Run

```
# Command
bash ./scripts/docker-run.sh "{DOCKER_HUB_REPOSITORY_TAGNAME}" "WECOM_ROBOT_CHECK={WECOM_ROBOT_CHECK}"

# Example
bash ./scripts/docker-run.sh "docker.io/mazeyqian/go-gin-gee:v20230427102021-api" "WECOM_ROBOT_CHECK=b2d57746-7146-44f2-8207-86cb0ca832be"
```

## Contributing

```
# All Dependences
go mod download

# Add
go get github.com/example/name
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn
```

It's necessary to run the command `go run scripts/init/main.go` when serving the project first.

```
# Serve
go run cmd/api/main.go -configpath="data/config.dev.json"

# Restart
# cd /web/go-gin-gee
go run scripts/restart/main.go
```

Visit: `http://127.0.0.1:3000/api/ping`.

```
pong/v1.0.0/2022-09-29 04:52:43
```

## References

1. [Gin Web Framework](https://github.com/gin-gonic/gin)
2. [lo](https://github.com/samber/lo)
3. [project-layout](https://github.com/golang-standards/project-layout)
4. [script](https://github.com/bitfield/script)
