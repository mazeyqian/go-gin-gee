package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// 文章导航 相关文章
	ignoreTitles := []string{
		"文章导航",
		"相关文章",
	}
	// Visited URLs
	visitedURLs := make(map[string]bool)

	// colly
	// 创建一个新的 Colly Collector
	c := colly.NewCollector(
		colly.AllowedDomains("blog.mazey.net"), // 限制爬取的域名
	)

	// 找到每个 `<h2>` 标签并打印内容
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		thatTitle := e.Text
		// 忽略标题
		for _, title := range ignoreTitles {
			if thatTitle == title {
				return
			}
		}
		// Ignore the title including "条评论"
		if len(thatTitle) > 6 && strings.Contains(thatTitle, "条评论") {
			fmt.Println("Ignore title:", thatTitle)
			return
		}
		fmt.Println("Title found:", thatTitle)
	})

	// 错误处理
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error occurred:", err)
	})

	// 处理分页
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Visit link found on page
		if visitedURLs[link] {
			return
		}
		// Handle the link end wiht .html
		if len(link) < 5 || link[len(link)-5:] != ".html" {
			fmt.Println("Ignore link:", link)
			return
		}
		fmt.Println("Next page found:", link)
		visitedURLs[link] = true
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// 开始爬取
	err := c.Visit("https://blog.mazey.net/4899.html") // 替换为目标网站 URL
	if err != nil {
		log.Fatal(err)
	}
}
