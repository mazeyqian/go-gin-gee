package main

import (
	"os"

	"github.com/mazeyqian/go-gin-gee/internal/startup"
)

func main() {
	os.Setenv("TZ", "Asia/Shanghai")
	startup.Run("data/startup-config.json", "json")
}
