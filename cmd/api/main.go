package main

import (
	"os"

	_ "github.com/mazeyqian/go-gin-gee/docs"
	"github.com/mazeyqian/go-gin-gee/internal/api"
	"github.com/mazeyqian/go-gin-gee/schedules"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @contact.name Antonio Paya Gonzalez
// @contact.url http://antoniopg.tk
// @contact.email antonioalfa22@gmail.com

// @license.name MIT
// @license.url https://github.com/mazeyqian/go-gin-gee/blob/master/LICENSE

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// https://www.zeitverschiebung.net/en/timezone/asia--shanghai
	os.Setenv("TZ", "Asia/Shanghai")
	schedules.RunCheck()
	api.Run("data/config.json", "json")
}
