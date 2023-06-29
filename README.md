# go-gin-gee

Gee is a project that provides several services for everyday work. The project is based on Gin [1], and follows the ProjectLayout [3] structure. In addition, some daily scripts in the folder `scripts` depend on Script [4], which can be used by the command `run` directly.

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

3\. Convert TypeDoc Comments to Markdown.

```
go run scripts/convert-typedoc-to-markdown/main.go
```

Usage: [English](https://github.com/mazeyqian/go-gin-gee/releases/tag/v1.2.0) | [简体中文](https://blog.mazey.net/3494.html#%E6%B3%A8%E9%87%8A%E8%BD%AC_Markdown)

4\. Convert Markdown to TypeDoc Comments.

```
go run scripts/convert-markdown-to-typedoc/main.go
```

Usage: [English](https://github.com/mazeyqian/go-gin-gee/releases/tag/v1.3.0) | [简体中文](https://blog.mazey.net/3494.html#Markdown_%E8%BD%AC%E6%B3%A8%E9%87%8A)

5\. Transfer Apple note table to Markdown table.

```
go run scripts/transfer-notes-to-md-table/main.go
```

More in folder `scripts`.

## API Examples

The base URL for this API is an environment variate `${BASE_URL}`, such as `https://example.com/path`.

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
curl --location --request POST '${BASE_URL}/api/gee/generate-short-link' \
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
    "tiny_link": "${BASE_URL}/t/b"
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
curl --location --request POST '${BASE_URL}/api/gee/create-alias2data' \
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
curl --location '${BASE_URL}/api/gee/get-data-by-alias?alias=alias%20example'
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

Environment Variates:

- `${WECOM_ROBOT_CHECK}`: WeCom Robot Key.
- `${BASE_URL}`: The Base URL for this Service.

### Supervisor

```
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api-linux-amd64 -configpath="/web/go-gin-gee/data/config.json"
autostart=true
autorestart=true
environment=WECOM_ROBOT_CHECK="b2lsjd46-7146-4nv2-8767-86cb0cncjdbe",BASE_URL="https://example.com/path"
```

### Docker

#### Build

Run `bash ./scripts/docker-build.sh -h` to see the help message.

```
Usage: docker-build.sh [OPTIONS] [ENV_VARS...]
Build and run a Docker container for the go-gin-gee API.

Options:
  -r, --run     Run the Docker container after building (default)
  -b, --build   Build the Docker image but do not run it
  -h, --help    Print this help message and exit

Environment variables:
  Any additional arguments passed to the script will be passed as environment variables to the Docker container.
```

**Usage**

`${RUN_FLAG}` is optional, default is `-r`("RUN"). `${WECOM_ROBOT_CHECK}` is optional. If you don't want to send the message to WeCom Robot, just remove it. `${BASE_URL}` is required. It's the Base URL for this Service.

```
bash ./scripts/docker-build.sh ${RUN_FLAG} "WECOM_ROBOT_CHECK=${WECOM_ROBOT_CHECK}" "BASE_URL=${BASE_URL}"
```

**Examples**

Example 1: Build and Push

```
bash ./scripts/docker-build.sh -b "WECOM_ROBOT_CHECK=b2lsjd46-7146-4nv2-8767-86cb0cncjdbe" "BASE_URL=https://example.com/path"
```

Example 2: Build and Run

```
bash ./scripts/docker-build.sh -r "WECOM_ROBOT_CHECK=b2lsjd46-7146-4nv2-8767-86cb0cncjdbe" "BASE_URL=https://example.com/path"
```

#### Run

Run `bash ./scripts/docker-run.sh -h` to see the help message.

```
Usage: docker-run.sh [OPTIONS] IMAGE_TAG [ENV_VARS...]
Run a Docker container from the specified IMAGE_TAG with the specified environment variables.

Options:
  -h, --help    Print this help message and exit

Environment variables:
  Any additional arguments passed to the script will be passed as environment variables to the Docker container.

Note:
  The first argument (IMAGE_TAG) must be the tag name of the Docker image to run.
```

Find the latest image tag name: [Tags](https://hub.docker.com/repository/docker/mazeyqian/go-gin-gee/tags?page=1&ordering=last_updated)

**Usage**

```
bash ./scripts/docker-run.sh "${DOCKER_HUB_REPOSITORY_TAGNAME}" "WECOM_ROBOT_CHECK=${WECOM_ROBOT_CHECK}" "BASE_URL=${BASE_URL}"
```

**Example**

```
bash ./scripts/docker-run.sh "docker.io/mazeyqian/go-gin-gee:v20230615221222-api" "WECOM_ROBOT_CHECK=b2lsjd46-7146-4nv2-8767-86cb0cncjdbe" "BASE_URL=https://example.com/path"
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
