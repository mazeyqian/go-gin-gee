package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func ConvertShanghaiToUTC(shanghaiTime string) (string, error) {
	shanghaiHour, err := strconv.Atoi(shanghaiTime[:2])
	if err != nil {
		return "", err
	}
	shanghaiMinute, err := strconv.Atoi(shanghaiTime[3:])
	if err != nil {
		return "", err
	}
	shanghaiTotalMinutes := shanghaiHour*60 + shanghaiMinute
	utcTotalMinutes := (shanghaiTotalMinutes - 480 + 1440) % 1440
	utcHour := utcTotalMinutes / 60
	utcMinute := utcTotalMinutes % 60
	utcTime := fmt.Sprintf("%02d:%02d", utcHour, utcMinute)
	return utcTime, nil
}

func RunCheck() {
	s := persistence.GetRobotRepository()
	// https://github.com/go-co-op/gocron
	// https://pkg.go.dev/time#Location
	// shanghai, _ := time.LoadLocation("Asia/Shanghai")
	UTC, _ := time.LoadLocation("UTC")
	ss := gocron.NewScheduler(UTC)
	// shTimeHour := 10
	// shTimeMinute := "00"
	// everyDayAtStr := fmt.Sprintf("%d:%s", shTimeHour-8, shTimeMinute)
	// Create a function to convert Asia/Shanghai TimeZone to UTC TimeZone.
	// Get a given Asia/Shanghai TimeZone string, such as "10:05", "04:01".
	// Return an UTC TimeZone string, such as "02:05", "20:01".
	everyDayAtStr, _ := ConvertShanghaiToUTC("10:00")
	log.Println("UTC everyDayAtStr:", everyDayAtStr)
	ss.Every(1).Day().At(everyDayAtStr).Do(s.ClearCheckResult)
	ss.StartAsync()
}
