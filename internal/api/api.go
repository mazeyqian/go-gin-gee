package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/internal/api/router"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/config"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/db"
)

func setConfiguration() {
	config.Setup()
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run() {
	setConfiguration()
	conf := config.GetConfig()
	log.Println("Config:", conf)
	// log.Println("conf.sites:", conf.Data.Sites, len(conf.Data.Sites))
	web := router.Setup()
	fmt.Println("Go API Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
