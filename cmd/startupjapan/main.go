package main

import (
	"os"

	"github.com/mazeyqian/go-gin-gee/internal/startupjapan"
)

func main() {
	os.Setenv("TZ", "Asia/Shanghai")
	startupjapan.Run("data/startup-config.json", "json")
}
