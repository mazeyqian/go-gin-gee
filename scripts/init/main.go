package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main() {
	fmt.Println("Init...")
	script.ListFiles("./asset").ExecForEach("cp -R {{.}} .").Stdout()
}
