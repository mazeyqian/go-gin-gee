package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main() {
	fmt.Println("I am scripts.")
	// script.Exec("ping 127.0.0.1").Stdout()
	script.FindFiles(".git").Stdout()
	// script.FindFiles("*.go").Stdout()
	// ExecForEach("gofmt -w {{ . }}")
}
