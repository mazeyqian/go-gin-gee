## 安装

确保你的 Go 环境已经安装，然后运行以下命令安装 Colly：

```bash
go get -u github.com/gocolly/colly/v2
```

## 基础用法

使用 Colly 爬取一个网站的标题。

```go
package main

import (
 "fmt"
 "log"

 "github.com/gocolly/colly/v2"
)

func main() {
 // 创建一个新的 Colly Collector
 c := colly.NewCollector(
  colly.AllowedDomains("example.com"), // 限制爬取的域名
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
 err := c.Visit("https://example.com") // 替换为目标网站 URL
 if err != nil {
  log.Fatal(err)
 }
}
```

输出结果：

```plain
Title found: Example Domain
```

## 常见回调函数

Colly 提供了许多回调函数，以下是最常用的几个：

### `OnHTML`

用于处理 HTML 元素，提取内容。

```go
c.OnHTML("div.article", func(e *colly.HTMLElement) {
 fmt.Println("Article found:", e.Text)
})
```

### `OnRequest`

在每次发送请求时触发，可以用来打印请求信息或修改请求头。

```go
c.OnRequest(func(r *colly.Request) {
 fmt.Println("Visiting:", r.URL)
 r.Headers.Set("User-Agent", "MyCustomUserAgent")
})
```

### `OnResponse`

在接收到响应时触发，可以用来处理原始的响应数据。

```go
c.OnResponse(func(r *colly.Response) {
 fmt.Println("Response received:", string(r.Body))
})
```

### `OnError`

在请求失败时触发。

```go
c.OnError(func(r *colly.Response, err error) {
 fmt.Println("Request failed:", r.Request.URL, "Error:", err)
})
```

## 处理多个页面

爬取所有页面需要递归访问下一页的链接，以下是一个分页爬取的示例：

### 示例代码

```go
package main

import (
 "fmt"
 "log"

 "github.com/gocolly/colly/v2"
)

func main() {
 c := colly.NewCollector()

 // 提取文章标题
 c.OnHTML(".article-title", func(e *colly.HTMLElement) {
  fmt.Println("Article Title:", e.Text)
 })

 // 处理分页
 c.OnHTML("a.next", func(e *colly.HTMLElement) {
  nextPage := e.Attr("href")
  fmt.Println("Next page found:", nextPage)
  c.Visit(e.Request.AbsoluteURL(nextPage))
 })

 // 错误处理
 c.OnError(func(r *colly.Response, err error) {
  log.Println("Error:", err)
 })

 // 开始爬取
 c.Visit("https://example.com/articles")
}
```

## 防反爬技巧

许多网站会有反爬机制，以下是一些常见的防反爬策略：

### 设置 User-Agent

模拟浏览器的请求头：

```go
c.OnRequest(func(r *colly.Request) {
 r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
})
```

#### **2. 添加请求延迟**

设置每次请求的间隔时间，避免频繁访问：

```go
c.Limit(&colly.LimitRule{
 DomainGlob:  "*example.com",
 Delay:       2 * time.Second,
 Parallelism: 1, // 单线程爬取
})
```

#### **3. 使用代理**

通过代理隐藏真实 IP：

```go
c.SetProxy("http://proxy.example.com:8080")
```

#### **4. 随机 User-Agent**

使用 `github.com/corpix/uarand` 库，随机生成 User-Agent：

```bash
go get github.com/corpix/uarand
```

```go
import "github.com/corpix/uarand"

c.OnRequest(func(r *colly.Request) {
 r.Headers.Set("User-Agent", uarand.GetRandom())
})
```

---

### **7. 数据存储**

爬取的数据可以存储到文件或数据库中：

#### **存储到文件**

使用 Go 的 `os` 包将数据写入文件：

```go
import (
 "os"
)

file, err := os.Create("output.txt")
if err != nil {
 log.Fatal(err)
}
defer file.Close()

file.WriteString("Your data here")
```

#### **存储到数据库**

可以使用 Go 的数据库库（如 `gorm`）将数据存储到数据库中。以下是将数据存储到 SQLite 的示例：

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

```go
import (
 "gorm.io/driver/sqlite"
 "gorm.io/gorm"
)

type Article struct {
 ID    uint   `gorm:"primaryKey"`
 Title string
}

db, err := gorm.Open(sqlite.Open("articles.db"), &gorm.Config{})
if err != nil {
 log.Fatal(err)
}

db.AutoMigrate(&Article{})

// 保存数据
db.Create(&Article{Title: "Example Article"})
```

---

### **8. 处理动态内容**

如果目标网站是通过 JavaScript 动态加载内容（如 SPA 网站），`Colly` 无法直接获取这些数据。这时可以使用以下方法：

#### **1. 使用 API**

检查网站的网络请求（浏览器开发者工具 -> Network），直接调用 API 获取数据。

#### **2. 使用 `chromedp`**

`chromedp` 是一个基于 Chrome 的浏览器自动化工具，可以渲染动态页面并提取内容。

安装：

```bash
go get -u github.com/chromedp/chromedp
```

示例代码：

```go
package main

import (
 "context"
 "fmt"
 "log"

 "github.com/chromedp/chromedp"
)

func main() {
 ctx, cancel := chromedp.NewContext(context.Background())
 defer cancel()

 var result string
 err := chromedp.Run(ctx,
  chromedp.Navigate("https://example.com"),
  chromedp.Text(".dynamic-content", &result, chromedp.NodeVisible),
 )
 if err != nil {
  log.Fatal(err)
 }

 fmt.Println("Dynamic Content:", result)
}
```

---

### **9. 学习资源**

- **官方文档**：[Colly Documentation](https://pkg.go.dev/github.com/gocolly/colly/v2)
- **GitHub 示例**：[Colly GitHub Repository](https://github.com/gocolly/colly)
- **实践项目**：尝试爬取你感兴趣的网站，逐步增加功能。

---

如果你有任何问题或需要帮助实现特定功能，请随时告诉我！ 😊
