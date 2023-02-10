package main

import (
	"os"

	"github.com/mazeyqian/go-gin-gee/internal/startupnode"
)

func main() {
	os.Setenv("TZ", "Asia/Shanghai")
	startupnode.Run("data/startup-config.json", "json")
}
