package main

import (
	"fmt"
	"os/exec"

	"github.com/bitfield/script"
)

func main() {
	// fmt.Println("I am scripts.")
	// script.Exec("ping 127.0.0.1").Stdout()
	// https://bitfieldconsulting.com/golang/scripting
	// cmdLines := "echo 233 && pwd"
	// cmdLines += " && echo {{.}}"
	// cmdLines += " && echo 22"
	// cmdLines += " && echo ---"
	// script.ListFiles("/Users/mazey/Web/Mazey/*/.git/config").ExecForEach(`/bin/sh -c echo "222222"; echo "333"; echo "444";`).Stdout()
	script.ListFiles("/Users/mazey/Web/Mazey/*/.git").FilterLine(func(s string) string {
		// cmdLines := "echo " + s + " - begin -;"
		cmdLines := "echo - begin -;"
		cmdLines += fmt.Sprintf("echo Path: %s;", s) // https://pkg.go.dev/fmt#Sprintf
		cmdLines += fmt.Sprintf("cd %s;", s)         // "cd " + s + ";"
		// cmdLines += "pwd;"
		cmdLines += `git config user.email "mazey@mazey.net";`
		cmdLines += `git config user.name "Mazey Chu";`
		cmdLines += "echo - end -;"
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Err:", err)
		}
		fmt.Printf("%s", result)
		// script.Exec("(echo 222 && echo 333)").Stdout()
		// fmt.Println("print " + s)
		// cmdLines1 := `git config user.email "mazey@mazey.net"`
		// cmdLines2 := `git config user.name "Mazey Chu";`
		// script.Exec("cd " + s).Stdout()
		// script.Exec("pwd").Stdout()
		// script.Exec(cmdLines1).Stdout()
		// script.Exec(cmdLines2).Stdout()
		return s
	}).Stdout()
	// script.FindFiles("*.go").Stdout()
	// ExecForEach("gofmt -w {{ . }}")

	// cmdSeq := "cd ~; pwd"
	// cmd := exec.Command("/bin/sh", "-c", cmdSeq)
	// result, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Printf("Output: %s", result)
}
