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
// go run scripts/batch-git-pull/main.go -path="/Users/mazey/Web/Rabbit" -projects="placeholder"
// path required
// projects optional
func main() {
	log.Println("Git pull...")
	placeholder := "unknown"
	// https://gobyexample.com/command-line-flags
	projectPath := flag.String("path", placeholder, "folder of projects")
	assignedProjects := flag.String("projects", ".", "assigned projects")
	flag.Parse()
	log.Println("projectPath:", *projectPath)
	log.Println("assignedProjects:", *assignedProjects)
	if *projectPath == placeholder {
		log.Panicln("path is required")
	}
	projects := []string{
		"placeholder",
	}
	regexStr := "^.+("
	for _, v := range projects {
		regexStr += fmt.Sprintf("%s|", v)
	}
	// Example: ^.+(placeholder|.)$
	regexStr += fmt.Sprintf("%s)\\/\\.git$", *assignedProjects)
	// log.Println("regexStr:", regexStr)
	regex := regexp.MustCompile(regexStr)
	script.ListFiles(fmt.Sprintf("%s/*/.git", *projectPath)).MatchRegexp(regex).FilterLine(func(s string) string {
		cmdLines := constants.ScriptStartMsg
		cmdLines += fmt.Sprintf("echo Path: %s;", s)
		cmdLines += fmt.Sprintf("cd %s;", s)
		// Control the branch: cmdLines += `git checkout master;`
		cmdLines += `cd ../;`
		cmdLines += `git pull;`
		cmdLines += constants.ScriptEndMsg
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("error:", err)
		}
		log.Printf("result: %s", result)
		return ""
	}).Stdout()
}
