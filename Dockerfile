FROM golang:1.19.5-alpine AS build_base

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git  git gcc g++

# time: missing Location in call to Time.In
# https://medium.com/freethreads/panic-time-missing-location-in-call-to-date-89d171811d3
# RUN apk --no-cache add tzdata

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./cmd/api/main.go

# Start fresh from a smaller image
# https://github.com/docker-library/golang/blob/8e04c39d2ce4466162418245c8b1178951021321/1.19/alpine3.17/Dockerfile
FROM alpine:3.17
RUN apk add ca-certificates
RUN apk --no-cache add tzdata

WORKDIR /app

COPY --from=build_base /src/out/app /app/api
COPY --from=build_base /src/data /app/data

RUN chmod +x api

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
ENTRYPOINT ./api
