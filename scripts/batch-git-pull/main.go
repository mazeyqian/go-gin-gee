package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/bitfield/script"
)

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
	// projectPath := "/web/i.mazey.net"
	projects := []string{
		// "go-gin-gee",
		// "mazey",
		// "tool",
		"placeholder",
	}
	regexStr := "^.+("
	for _, v := range projects {
		regexStr += fmt.Sprintf("%s|", v)
	}
	regexStr += fmt.Sprintf("%s)$", *assignedProjects) // placeholder
	log.Println("regexStr:", regexStr)
	regex := regexp.MustCompile(regexStr) // "^.+(json-to-resume|mazey-server)$")
	script.ListFiles(*projectPath).MatchRegexp(regex).FilterLine(func(s string) string {
		cmdLines := "echo - - begin - -;"
		cmdLines += "echo ;"
		cmdLines += fmt.Sprintf("echo Path: %s;", s)
		cmdLines += fmt.Sprintf("cd %s;", s)
		// cmdLines += `git checkout master;`
		cmdLines += `git pull;`
		cmdLines += "echo ;"
		cmdLines += "echo - - end - - - - - - - - - - - - - - - - -;"
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("error:", err)
		}
		log.Printf("result: %s", result)
		return s
	}).Stdout()
}
