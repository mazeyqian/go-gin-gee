package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bitfield/script"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
)

// Example: go run scripts/convert-typedoc-to-markdown/main.go
func main() {
	log.Println(constants.StartMsgStr)
	alias := "td2md"
	// endLine := " */"
	// td2md.js td2md.md
	iFilename := fmt.Sprintf("./data/%s.js", alias)
	oFilename := fmt.Sprintf("./data/%s.md", alias)
	// index := 0
	// hasZH := false
	// https://gobyexample.com/regular-expressions
	// rMatchZH, _ := regexp.Compile("^<!-- (.+) -->$")
	// rMatchCommentBoth, _ := regexp.Compile("^<!-- (.+) -->$")
	rMatchCommentStart, _ := regexp.Compile(`^ *\*?\s?(.*)$`)
	rMatchCommentStartEN, _ := regexp.Compile(`^ *\*?\s?EN:\s(.*)$`)
	placeholderChar := "#loading..."
	script.File(iFilename).Match(" *").FilterLine(func(s string) string {
		retStr := s
		if strings.Contains(s, "/**") {
			// log.Println("First Line: ", retStr)
			retStr = placeholderChar // "#p"
		} else if strings.Contains(s, " */") {
			// log.Println("Last Line: ", retStr)
			retStr = placeholderChar // "#p"
		} else if strings.Contains(s, " * @") {
			// log.Println("Ignore Line: ", retStr)
			retStr = placeholderChar // "#p"
		} else if strings.Contains(s, " * ZH:") {
			// log.Println("Ignore Line: ", retStr)
			retStr = placeholderChar // "#p"
		} else if strings.Contains(s, " * EN:") {
			// log.Println("Ignore Line: ", retStr)
			retStr = rMatchCommentStartEN.FindStringSubmatch(retStr)[1]
		} else if strings.Contains(s, " *") || strings.Contains(s, " * ") {
			retStr = rMatchCommentStart.FindStringSubmatch(retStr)[1]
			// log.Println("Normal Line: ", retStr)
		} else {
			log.Println("Error: ", retStr)
		}
		log.Println(retStr)
		return retStr
		// Filter #p
	}).Reject(placeholderChar).WriteFile(oFilename)
	// https://pkg.go.dev/github.com/bitfield/script#Pipe.AppendFile
	log.Println(constants.EndMsgStr)

	// Read the output file and remove extra empty lines
	re2 := regexp.MustCompile(`\n{3,}`)
	re3 := regexp.MustCompile(`\n+$`)
	output, err := os.ReadFile(oFilename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	outputStr := string(output)
	outputStr = re2.ReplaceAllString(outputStr, "\n\n")
	outputStr = re3.ReplaceAllString(outputStr, "")
	output = []byte(outputStr)

	// Write the trimmed output back to the output file
	err = os.WriteFile(oFilename, output, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("All done.")
}
