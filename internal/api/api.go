package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/internal/api/router"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/config"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/db"
)

func setConfiguration() { // configPath string, configType string) {
	config.Setup() // configPath, configType)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run() { // configPath string, configType string) {
	// if configPath == "" {
	// 	configPath = "data/config.yml"
	// }
	// if configType == "" {
	// 	configType = "yaml"
	// }
	setConfiguration() // configPath, configType)
	conf := config.GetConfig()
	log.Println("conf:", conf)
	log.Println("conf.sites:", conf.Data.Sites, len(conf.Data.Sites))
	web := router.Setup()
	fmt.Println("Go API Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
