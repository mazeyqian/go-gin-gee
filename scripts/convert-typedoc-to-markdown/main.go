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

// func main() {
// 	log.Println(constants.StartMsgStr)
// 	// Open the input file
// 	inputFile, err := os.Open("./data/td2md.js")
// 	if err != nil {
// 		fmt.Println("Error opening input file:", err)
// 		return
// 	}
// 	defer inputFile.Close()

// 	// Create the output file
// 	outputFile, err := os.Create("./data/td2md.md")
// 	if err != nil {
// 		fmt.Println("Error creating output file:", err)
// 		return
// 	}
// 	defer outputFile.Close()

// 	// Copy the input to the output, removing the " * " prefix from each line
// 	var prefix = []byte(" * ")
// 	var buffer = make([]byte, 1024)
// 	for {
// 		n, err := inputFile.Read(buffer)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("Error reading input file:", err)
// 			return
// 		}
// 		outputFile.Write([]byte(strings.ReplaceAll(string(buffer[:n]), string(prefix), "")))
// 	}

// 	fmt.Println("Conversion complete!")
// }

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
	// rMatchCommentEnd, _ := regexp.Compile("^(.+) -->$")
	// zhStartStr := " * ZH:"
	// Determine the intel evironment.
	// script.File(iFilename).FilterLine(func(s string) string {
	// 	if strings.Contains(s, zhStartStr) {
	// 		hasZH = true
	// 		log.Println("ZH exist")
	// 	}
	// 	return constants.RunningMsg
	// }).Stdout()
	// Add comments.
	// Filter not #p
	// error parsing regexp: invalid or unsupported Perl syntax: `(?!`
	// re := regexp.MustCompile(`^(?!#p)`)
	// EN:
	// ZH:
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
		// First Line
		// if index == 0 {
		// 	if hasZH {
		// 		retStr = "/**\n * EN: " + retStr
		// 	} else {
		// 		retStr = "/**\n * " + retStr
		// 	}
		// } else {
		// 	// ZH
		// 	if strings.Contains(s, zhStartStr) {
		// 		retStr = " * " + rMatchZH.FindStringSubmatch(retStr)[1]
		// 	} else if strings.Contains(s, "<!-- ") && strings.Contains(s, " -->") {
		// 		retStr = " * " + rMatchCommentBoth.FindStringSubmatch(retStr)[1]
		// 	} else if strings.Contains(s, "<!-- ") {
		// 		retStr = " * " + rMatchCommentStart.FindStringSubmatch(retStr)[1]
		// 	} else if strings.Contains(s, " -->") {
		// 		retStr = " * " + rMatchCommentEnd.FindStringSubmatch(retStr)[1]
		// 	} else { // Normal Lines
		// 		retStr = " * " + retStr
		// 	}
		// }
		// index++
		log.Println(retStr)
		return retStr
		// Filter #p
	}).Reject(placeholderChar).WriteFile(oFilename)
	// https://pkg.go.dev/github.com/bitfield/script#Pipe.AppendFile
	// script.Echo(endLine).AppendFile(oFilename)
	// log.Println(endLine)
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
