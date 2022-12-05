package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/bitfield/script"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
)

// Examples:
// go run scripts/batch-git-pull/main.go -path="/Users/mazey/Web/Mazey"
// go run scripts/batch-git-pull/main.go -path="/Users/mazey/Web/Bilibili" -projects="placeholder"
// path required
// projects optional
func main() {
	log.Println("Git pull...")
	// https://gobyexample.com/command-line-flags
	projectPath := flag.String("path", "/web/i.mazey.net", "folder of projects")
	assignedProjects := flag.String("projects", ".", "assigned projects")
	flag.Parse()
	log.Println("projectPath:", *projectPath)
	log.Println("assignedProjects:", *assignedProjects)
	projects := []string{
		"placeholder",
	}
	regexStr := "^.+("
	for _, v := range projects {
		regexStr += fmt.Sprintf("%s|", v)
	}
	// Example: ^.+(placeholder|.)$
	regexStr += fmt.Sprintf("%s)$", *assignedProjects)
	// Example: /^(.+，)?([^（），]+)(.+)?$/
	// exclude .DS_Store
	if *assignedProjects == "." {
		regexStr = "^.+[^._ae]$"
	}
	log.Println("regexStr:", regexStr)
	regex := regexp.MustCompile(regexStr)
	script.ListFiles(*projectPath).MatchRegexp(regex).FilterLine(func(s string) string {
		cmdLines := constants.ScriptStartMsg // "echo - - begin - -;"
		cmdLines += fmt.Sprintf("echo Path: %s;", s)
		cmdLines += fmt.Sprintf("cd %s;", s)
		// Control the branch: cmdLines += `git checkout master;`
		cmdLines += `git pull;`
		cmdLines += constants.ScriptEndMsg // "echo - - end - - - - - - - - - - - - - - - - -;"
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("error:", err)
		}
		log.Printf("result: %s", result)
		return ""
	}).Stdout()
}
