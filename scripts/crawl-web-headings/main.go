package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// colly
	// 创建一个新的 Colly Collector
	c := colly.NewCollector(
		colly.AllowedDomains("blog.mazey.net"), // 限制爬取的域名
	)

	// 找到每个 `<h1>` 标签并打印内容
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Title found:", e.Text)
	})

	// 错误处理
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error occurred:", err)
	})

	// 开始爬取
	err := c.Visit("https://blog.mazey.net/2956.html") // 替换为目标网站 URL
	if err != nil {
		log.Fatal(err)
	}
}
