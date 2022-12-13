package main

import (
	"fmt"
	"log"

	"github.com/bitfield/script"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
)

func main() {
	log.Println(constants.StartMsgStr)
	alias := "markdown"
	iFilename := fmt.Sprintf("./data/%s-for-typedoc.md", alias)
	oFilename := fmt.Sprintf("./data/%s-for-typedoc.js", alias)
	index := 0
	script.File(iFilename).FilterLine(func(s string) string {
		retStr := s
		if index == 0 {
			retStr = "/**\n * " + retStr
		} else {
			retStr = " * " + retStr
		}
		index++
		return retStr
	}).WriteFile(oFilename)
	// https://pkg.go.dev/github.com/bitfield/script#Pipe.AppendFile
	script.Echo(" */").AppendFile(oFilename)
	log.Println(constants.EndMsgStr)
}
