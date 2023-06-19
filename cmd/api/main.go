package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/mazeyqian/go-gin-gee/docs"
	"github.com/mazeyqian/go-gin-gee/internal/api"
	"github.com/mazeyqian/go-gin-gee/internal/api/controllers"
)

// @Golang API
// @version 1.0
// @description API in Golang with Gin Framework

// @contact.name Mazey Chu
// @contact.url https://blog.mazey.net
// @contact.email mazeyqian@gmail.com

// @license.name MIT
// @license.url https://github.com/mazeyqian/go-gin-gee/blob/master/LICENSE

// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configPath := flag.String("configpath", "data/config.json", "path of configuration")
	flag.Parse()
	log.Println("configPath:", *configPath)
	// ENV
	// Robot
	// MacOS, export WECOM_ROBOT_CHECK="b2lsjd46-7146-4nv2-8767-86cb0cncjdbe"
	// https://knowhowcommunity.org/how-to-set-environment-variables-mac/
	WECOM_ROBOT_CHECK := os.Getenv("WECOM_ROBOT_CHECK")
	log.Println("WECOM_ROBOT_CHECK:", WECOM_ROBOT_CHECK)
	// https://www.zeitverschiebung.net/en/timezone/asia--shanghai
	// os.Setenv("TZ", "Asia/Shanghai")
	// Set the timezone to UTC
	os.Setenv("TZ", "UTC")
	controllers.RunCheck()
	api.Run(*configPath, "json")
}
