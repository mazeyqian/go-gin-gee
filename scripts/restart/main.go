package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Restart...")

	cmdLines := "echo - begin -;"
	cmdLines += `cd /web/go-gin-gee;`
	cmdLines += `git checkout master;`
	cmdLines += `git pull;`
	cmdLines += `GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go;`
	// cmdLines += `supervisorctl restart api;`
	cmdLines += `systemctl restart supervisord;`
	cmdLines += `supervisorctl status;`
	cmdLines += "echo - end -;"
	cmd := exec.Command("/bin/sh", "-c", cmdLines)
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Printf("%s", result)
}
