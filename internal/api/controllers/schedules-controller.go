package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
)

func CheckSitesHealth(c *gin.Context) {
	s := persistence.GetRobotRepository()
	markdown, _ := s.ClearCheckResult()
	c.JSON(http.StatusOK, gin.H{"data": *markdown})
}

func RunCheck() {
	s := persistence.GetRobotRepository()
	// https://github.com/go-co-op/gocron
	// https://pkg.go.dev/time#Location
	shanghai, _ := time.LoadLocation("Asia/Shanghai")
	ss := gocron.NewScheduler(shanghai)
	ss.Every(1).Day().At("10:00").Do(s.ClearCheckResult)
	ss.StartAsync()
}
