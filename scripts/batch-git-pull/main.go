package main

import (
	"fmt"
	"os/exec"

	"github.com/bitfield/script"
)

func main() {
	fmt.Println("Change git user...")

	script.ListFiles("/Users/mazey/Web/Mazey").FilterLine(func(s string) string {
		cmdLines := "echo - begin -;"
		cmdLines += fmt.Sprintf("echo Path: %s;", s)
		cmdLines += fmt.Sprintf("cd %s;", s)
		cmdLines += `git checkout master;`
		cmdLines += `git pull;`
		cmdLines += "echo - end -;"
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Err:", err)
		}
		fmt.Printf("%s", result)
		return s
	}).Stdout()
}
