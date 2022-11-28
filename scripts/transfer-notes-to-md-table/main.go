package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	log.Println("- begin -")
	re := regexp.MustCompile("^.*\t.*\t.*$")
	script.File("/Users/mazey/Web/Mazey/go-gin-gee/data/s.txt").MatchRegexp(re).FilterLine(func(s string) string {
		ss := strings.Split(s, "\t")
		sss := ss[:3]
		if sss[0] == "Title" || sss[0] == "" || len(sss) != 3 {
			return "| - | - | - | # |"
		}
		retS := ""
		for _, v := range sss {
			ssss := fmt.Sprintf("| %s ", v)
			retS += ssss
		}
		retS += "| # |"
		s = retS
		return s
	}).WriteFile("./data/sw.txt")
	log.Println("- end -")
}
