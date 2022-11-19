# API
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go

# Startup
GOOS=darwin GOARCH=amd64 go build -o dist/startup-mac-darwin-amd64 cmd/startup/main.go

# Startup Node
GOOS=linux GOARCH=amd64 go build -o dist/startupnode-linux-amd64 cmd/startupnode/main.go