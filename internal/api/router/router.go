package router

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/internal/api/controllers"
	"github.com/mazeyqian/go-gin-gee/internal/api/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	if err := os.MkdirAll("log", 0755); err != nil {
		log.Println("mkdir err:", err)
	}
	f, err := os.Create("log/api.log")
	if err != nil {
		log.Println("create err:", err)
	}
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
	app.Use(middlewares.CORS())
	app.Use(middlewares.Logger())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// ================== Login Routes
	app.POST("/api/login", controllers.Login)
	app.POST("/api/login/add", controllers.CreateUser)
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ================== User Routes
	app.GET("/api/users", controllers.GetUsers)
	app.GET("/api/users/:id", controllers.GetUserById)
	app.POST("/api/users", controllers.CreateUser)
	app.PUT("/api/users/:id", controllers.UpdateUser)
	app.DELETE("/api/users/:id", controllers.DeleteUser)
	// ================== Tasks Routes
	app.GET("/api/tasks/:id", controllers.GetTaskById)
	app.GET("/api/tasks", controllers.GetTasks)
	app.POST("/api/tasks", controllers.CreateTask)
	app.PUT("/api/tasks/:id", controllers.UpdateTask)
	app.DELETE("/api/tasks/:id", controllers.DeleteTask)

	// Gin Examples - begin
	app.GET("/api/ping", controllers.Ping)
	app.GET("/api/Get-Custom-Struct", controllers.GetDataB)
	app.GET("/api/AsciiJSON", controllers.AsciiJSON)
	app.POST("/api/Bind-html-checkboxes", controllers.FormHandler)
	app.GET("/api/bind-query-or-post", controllers.StartPage)
	app.GET("/api/Middleware", controllers.Middleware0920)
	// app.LoadHTMLGlob("templates/*")
	app.GET("/api/index", controllers.Index0920)
	app.GET("/api/JSONP", controllers.JSONP0920)
	app.GET("/api/:name/:id", controllers.NameId0920)
	app.POST("/api/postform-parameters", controllers.PostformParameters)
	app.GET("/api/resty-ping", controllers.RestyPing)
	app.GET("/api/resty-upload", controllers.RestyUpload)
	// Grouping routes
	v2 := app.Group("/api/v2")
	{
		v2.GET("/bind-query-or-post", controllers.StartPage)
		v2.POST("/Bind-html-checkboxes", controllers.FormHandler)
	}
	// Gin Examples - end

	// Gee API - begin
	gee := app.Group("/api/gee")
	{
		gee.GET("/get-data-by-alias", controllers.GetDataByAlias)
		gee.POST("/create-alias2data", controllers.CreateAlias2data)
		gee.GET("/count-alias2data", controllers.CountAlias2data)
		gee.GET("/check", controllers.CheckSitesHealth)
		gee.GET("/query-short-link", controllers.GetTiny)
		gee.POST("/generate-short-link", controllers.CreateTiny)
		gee.GET("/get-tag-name", controllers.GetTag)
	}
	// Tiny
	app.GET("/t/:key", controllers.RedirectTiny)
	// Gee API - end

	return app
}
