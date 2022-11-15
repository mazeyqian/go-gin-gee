package main

import (
	"os"

	"github.com/mazeyqian/go-gin-gee/internal/startupnode"
)

func main() {
	os.Setenv("TZ", "Asia/Shanghai")
	startupnode.Run("data/startupnode-config.json", "json")
}
