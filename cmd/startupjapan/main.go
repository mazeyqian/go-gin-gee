package main

import (
	"flag"
	"log"
	"os"

	"github.com/mazeyqian/go-gin-gee/internal/startupjapan"
)

// Examples:
// go run cmd/startupjapan/main.go -projectpath "/Users/mazey/Web/Mazey/go-gin-gee"
// go run cmd/startupjapan/main.go -projectpath "/web/go-gin-gee"
func main() {
	projectPath := flag.String("projectpath", "/web/go-gin-gee", "path of project")
	flag.Parse()
	log.Println("projectPath:", *projectPath)
	os.Setenv("TZ", "Asia/Shanghai")
	startupjapan.Run("data/startup-config.json", "json", *projectPath)
}
