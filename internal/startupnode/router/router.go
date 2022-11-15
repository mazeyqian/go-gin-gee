package router

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/startupnode.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return app
}
