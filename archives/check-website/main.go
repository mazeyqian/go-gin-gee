package main

import (
	"fmt"

	"github.com/mazeyqian/go-gin-gee/archives/schedules"
)

func main() {
	fmt.Println("Check website...")
	schedules.Check()
}
