package main

import (
	"os"

	_ "github.com/mazeyqian/go-gin-gee/docs"
	"github.com/mazeyqian/go-gin-gee/internal/api"
	"github.com/mazeyqian/go-gin-gee/internal/api/controllers"
)

// @Golang API
// @version 1.0
// @description API in Golang with Gin Framework

// @contact.name Cheng
// @contact.url https://github.com/mazeyqian
// @contact.email mazeyqian@gmail.com

// @license.name MIT
// @license.url https://github.com/mazeyqian/go-gin-gee/blob/main/LICENSE

// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Set the timezone to UTC
	// https://www.zeitverschiebung.net/en/timezone/asia--shanghai
	os.Setenv("TZ", "UTC")
	// Run before the API starts
	controllers.RunCheck()
	api.Run()
}
