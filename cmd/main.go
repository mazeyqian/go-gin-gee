package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var db = make(map[string]string)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}

type PersonBindUrl struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
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

// resty client
var client = resty.New()

func setupRouter() *gin.Engine {
	db["mazey"] = "cherrie"
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Gin Examples - begin

	// Custom Middleware
	// https://gin-gonic.com/docs/examples/custom-middleware/
	r.Use(Logger())
	r.GET("/Middleware", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)

		c.JSON(200, gin.H{"testMiddleware": example})
	})

	// Ping test
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "pong")
	// })

	// Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
	// https://gin-gonic.com/docs/examples/ascii-json/
	// r.GET("/AsciiJSON", func(c *gin.Context) {
	// 	data := map[string]interface{}{
	// 		"lang": "GO语言",
	// 		"tag":  "<br>",
	// 	}

	// 	// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	// 	c.AsciiJSON(http.StatusOK, data)
	// })

	// Bind form-data request with custom struct
	// https://gin-gonic.com/docs/examples/bind-form-data-request-with-custom-struct/
	// r.GET("/Get-Custom-Struct", GetDataB)

	// Bind html checkboxes
	// https://gin-gonic.com/docs/examples/bind-html-checkbox/
	// r.POST("/Bind-html-checkboxes", formHandler)

	// Bind query string or post data
	// https://gin-gonic.com/docs/examples/bind-query-or-post/
	r.GET("/bind-query-or-post", startPage)

	// Bind Uri
	// https://gin-gonic.com/docs/examples/bind-uri/
	r.GET("/:name/:id", func(c *gin.Context) {
		var person PersonBindUrl
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Grouping routes
	v2 := r.Group("/v2")
	{
		v2.GET("/bind-query-or-post", startPage)
		v2.POST("/Bind-html-checkboxes", formHandler)
	}

	// HTML rendering
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// JSONP
	// https://gin-gonic.com/docs/examples/jsonp/
	// Wrong
	r.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// Map as querystring or postform parameters
	// https://gin-gonic.com/docs/examples/map-as-querystring-or-postform/
	r.POST("/postform-parameters", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		log.Printf("ids: %v; names: %v", ids, names)
		c.JSON(http.StatusOK, gin.H{"ids": ids, "names": names})
	})

	// Gin Examples - end

	// Gin Introduction - begin

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	// Gin Introduction - end

	// test - begin
	r.GET("/resty-ping", func(c *gin.Context) {
		// Create a Resty Client
		client := resty.New()
		resp, err := client.R().
			EnableTrace().
			Get("https://httpbin.org/get")
		// Explore response object
		log.Println("Response Info:")
		log.Println("  Error      :", err)
		log.Println("  Status Code:", resp.StatusCode())
		log.Println("  Status     :", resp.Status())
		log.Println("  Proto      :", resp.Proto())
		log.Println("  Time       :", resp.Time())
		log.Println("  Received At:", resp.ReceivedAt())
		log.Println("  Body       :\n", resp)
		log.Println()

		// Explore trace info
		fmt.Println("Request Trace Info:")
		ti := resp.Request.TraceInfo()
		fmt.Println("  DNSLookup     :", ti.DNSLookup)
		fmt.Println("  ConnTime      :", ti.ConnTime)
		fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
		fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
		fmt.Println("  ServerTime    :", ti.ServerTime)
		fmt.Println("  ResponseTime  :", ti.ResponseTime)
		fmt.Println("  TotalTime     :", ti.TotalTime)
		fmt.Println("  IsConnReused  :", ti.IsConnReused)
		fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
		fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
		fmt.Println("  RequestAttempt:", ti.RequestAttempt)
		fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	r.GET("/resty-upload", func(c *gin.Context) {
		// POST of raw bytes for file upload.
		fileBytes, _ := ioutil.ReadFile("./data/in.csv")

		// See we are not setting content-type header, since go-resty automatically detects Content-Type for you
		resp, err := client.R().
			SetHeader("Content-Type", "application/octet-stream").
			SetBody(bytes.NewReader(fileBytes)).
			SetContentLength(true). // Dropbox expects this value
			Post("http://localhost:3002/")

		// Explore response object
		log.Println("Response Info:")
		log.Println("  Error      :", err)
		log.Println("  Status Code:", resp.StatusCode())
		log.Println("  Status     :", resp.Status())
		log.Println("  Proto      :", resp.Proto())
		log.Println("  Time       :", resp.Time())
		log.Println("  Received At:", resp.ReceivedAt())
		log.Println("  Body       :\n", resp)
		log.Println()
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	// test - end

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8214")
}
