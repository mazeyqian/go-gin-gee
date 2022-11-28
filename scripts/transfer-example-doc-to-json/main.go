package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/bitfield/script"
)

func main() {
	log.Println("- begin -")
	re := regexp.MustCompile("^.+$")
	index := 0
	script.File("./data/example-plain.txt").MatchRegexp(re).FilterLine(func(s string) string {
		retS := ""
		retS += fmt.Sprintf("\"example%d\": '%s',", index, s)
		index++
		s = retS
		log.Println("New line:", s)
		return s
	}).WriteFile("./data/example-json.txt")
	log.Println("- end -")
}
