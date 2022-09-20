package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "Metodo no permitido"})
	}
}

// NoRouteHandler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "The processing function of the request route was not found"})
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("Consume Time: ", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("StatusCode: ", status)
	}
}
