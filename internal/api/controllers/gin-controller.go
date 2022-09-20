package controllers

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

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

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

func AsciiJSON(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

	// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func FormHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func StartPage(c *gin.Context) {
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

func Middleware0920(c *gin.Context) {
	example := c.MustGet("example").(string)

	// it would print: "12345"
	log.Println(example)

	c.JSON(200, gin.H{"testMiddleware": example})
}

func Index0920(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func JSONP0920(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	//callback is x
	// Will output  :   x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}

type PersonBindUrl struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func NameId0920(c *gin.Context) {
	var person PersonBindUrl
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}

func PostformParameters(c *gin.Context) {

	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	log.Printf("ids: %v; names: %v", ids, names)
	c.JSON(http.StatusOK, gin.H{"ids": ids, "names": names})
}

func RestyPing(c *gin.Context) {
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
}

func RestyUpload(c *gin.Context) {
	var client = resty.New()
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
}
