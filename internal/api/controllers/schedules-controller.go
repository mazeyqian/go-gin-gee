package controllers

import (
	"fmt"
	"log"
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
	// shanghai, _ := time.LoadLocation("Asia/Shanghai")
	UTC, _ := time.LoadLocation("UTC")
	ss := gocron.NewScheduler(UTC)
	shTimeHour := 10
	shTimeMinute := "00"
	everyDayAtStr := fmt.Sprintf("%d:%s", shTimeHour-8, shTimeMinute)
	log.Println("UTC everyDayAtStr:", everyDayAtStr)
	ss.Every(1).Day().At(everyDayAtStr).Do(s.ClearCheckResult)
	ss.StartAsync()
}
