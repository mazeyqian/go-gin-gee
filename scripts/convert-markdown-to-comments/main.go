package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bitfield/script"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
)

// Example: go run scripts/convert-markdown-to-comments/main.go
func main() {
	log.Println(constants.StartMsgStr)
	alias := "markdown"
	endLine := " */"
	// markdown-for-typedoc.md
	iFilename := fmt.Sprintf("./data/%s-for-typedoc.md", alias)
	oFilename := fmt.Sprintf("./data/%s-for-typedoc.js", alias)
	index := 0
	hasZH := false
	// https://gobyexample.com/regular-expressions
	rMatchZH, _ := regexp.Compile("^<!-- (.+) -->$")
	// Determine the intel evironment.
	script.File(iFilename).FilterLine(func(s string) string {
		if strings.Contains(s, "ZH:") {
			hasZH = true
			log.Println("ZH exist")
		}
		return constants.RunningMsg
	}).Stdout()
	// Add comments.
	script.File(iFilename).FilterLine(func(s string) string {
		retStr := s
		if index == 0 {
			if hasZH {
				retStr = "/**\n * EN: " + retStr
			} else {
				retStr = "/**\n * " + retStr
			}
		} else {
			if strings.Contains(s, "ZH:") {
				retStr = " * " + rMatchZH.FindStringSubmatch(retStr)[1]
			} else {
				retStr = " * " + retStr
			}
		}
		index++
		log.Println(retStr)
		return retStr
	}).WriteFile(oFilename)
	// https://pkg.go.dev/github.com/bitfield/script#Pipe.AppendFile
	script.Echo(endLine).AppendFile(oFilename)
	log.Println(endLine)
	log.Println(constants.EndMsgStr)
}
