package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "metodo no permitido"})
	}
}

// NoRouteHandler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "not found"})
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Request URL:", c.Request.URL)

		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Println("Consume Time:", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("StatusCode:", status)
	}
}
