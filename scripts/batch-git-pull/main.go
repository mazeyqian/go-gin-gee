package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/bitfield/script"
)

func main() {
	log.Println("Change git user...")
	projectPath := "/Users/mazey/Web/Mazey"
	projects := []string{
		"go-gin-gee",
		"json-to-resume",
		"mazey",
		"mazey-study",
		"mazey-server",
	}
	regexStr := "^.+("
	for _, v := range projects {
		regexStr += fmt.Sprintf("%s|", v)
	}
	regexStr += "placeholder)$"
	log.Println("regexStr:", regexStr)
	regex := regexp.MustCompile(regexStr) // "^.+(json-to-resume|mazey-server)$")
	script.ListFiles(projectPath).MatchRegexp(regex).FilterLine(func(s string) string {
		cmdLines := "echo - - begin - -;"
		cmdLines += "echo ;"
		cmdLines += fmt.Sprintf("echo Path: %s;", s)
		// cmdLines += fmt.Sprintf("cd %s;", s)
		// cmdLines += `git checkout master;`
		// cmdLines += `git pull;`
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
