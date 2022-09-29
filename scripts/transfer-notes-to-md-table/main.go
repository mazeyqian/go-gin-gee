package main

import (
	"fmt"
	"regexp"

	"github.com/bitfield/script"
)

func main() {
	fmt.Println("Init...")

	contents, _ := script.File("/Users/mazey/Web/Mazey/go-gin-gee/data/s.txt").String()

	fmt.Println(contents)

	fmt.Println("- begin -")

	re := regexp.MustCompile("^.*\t.*((?!\t).)*$")

	// script.File("/Users/mazey/Web/Mazey/go-gin-gee/data/s.txt").ExecForEach("echo -, {{.}}").Stdout()
	script.File("/Users/mazey/Web/Mazey/go-gin-gee/data/s.txt").MatchRegexp(re).ExecForEach("echo -, {{.}}").Stdout()

	fmt.Println("- end -")
}
