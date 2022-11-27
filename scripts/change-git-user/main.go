package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/bitfield/script"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
)

func main() {
	log.Println("Change git user...")
	/*
	 * Air Example: "/Volumes/ProjectX/Example"
	 * "example@example.net" "Example Na"
	 * Pro Mazey: "/Users/mazey/Web/Mazey"
	 * "mazey@mazey.net" "Mazey Chu"
	 */
	absolutePath := "/Users/mazey/Web/Mazey"
	userEmail := "mazey@mazey.net"
	userName := "Mazey Chu"
	// https://bitfieldconsulting.com/golang/scripting
	// https://pkg.go.dev/github.com/bitfield/script#ListFiles
	script.ListFiles(fmt.Sprintf("%s/*/.git", absolutePath)).FilterLine(func(s string) string {
		cmdLines := constants.ScriptStartMsg
		// https://pkg.go.dev/fmt#Sprintf
		cmdLines += fmt.Sprintf("echo Path: %s;", s)
		cmdLines += fmt.Sprintf("cd %s;", s)
		cmdLines += fmt.Sprintf(`git config user.email "%s";`, userEmail)
		cmdLines += fmt.Sprintf(`git config user.name "%s";`, userName)
		cmdLines += constants.ScriptEndMsg
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("error:", err)
		}
		log.Printf("%s", result)
		return ""
	}).Stdout()
}
