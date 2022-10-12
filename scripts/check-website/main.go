package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	wxworkbot "github.com/vimsucks/wxwork-bot-go"
)

type Sites struct {
	List map[string]SiteStatus
}

type SiteStatus struct {
	Name string
	Code int
}

func (s *Sites) getWebSiteStatus() ([]string, error) {
	// http://c.biancheng.net/view/32.html
	healthySites := []string{}
	client := resty.New()
	for url, status := range s.List {
		log.Println("url:", url)
		log.Println("code expect:", status.Code)
		resCode := 0
		resp, err := client.R().
			Get(url)
		if err != nil {
			// return 0, err
			log.Println("error:", err)
		} else {
			resCode = resp.StatusCode()
			log.Println("code get:", resp.StatusCode())
		}
		if status.Code == resCode {
			healthySites = append(healthySites, status.Name)
		}
	}
	// resp, err := client.R().
	// 	Get(url)
	// if err != nil {
	// 	return []string{"0"}, err
	// }
	return healthySites, nil
}

func main() {
	fmt.Println("Check website...")
	type Person struct {
		Name string `json:"name"`
		Id   string `json:"uuid"`
	}
	p := &Person{}
	client := resty.New()
	resp, err := client.R().
		// EnableTrace().
		SetResult(p).
		Get("https://feperf.com/api/mazeychu/413fbc97-4bed-3228-9f07-9141ba07c9f3")
	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :", resp)
	log.Println()
	log.Println("  Name:", p.Name)
	log.Println("  Id:", p.Id)

	ss := &Sites{}
	ss.List = make(map[string]SiteStatus)
	ss.List["https://blog.mazey.net/"] = SiteStatus{"博客首页", 200}
	ss.List[fmt.Sprintf("%s%d", "https://blog.mazey.net/?s=", time.Now().Unix())] = SiteStatus{"博客搜索页", 200}
	ss.List["https://tool.mazey.net/markdown/"] = SiteStatus{"Markdown", 200}
	healthySites, err := ss.getWebSiteStatus()
	if err != nil {
		log.Println("  Error      :", err)
	}
	log.Println("Healthy Sites:", healthySites)
	mdStr := "Health Check Result:\n"
	for _, siteName := range healthySites {
		mdStr += fmt.Sprintf("%s OK\n", siteName)
	}
	mdStr += fmt.Sprintf("%s%d", "Sum: ", len(healthySites))
	log.Println(mdStr)
	// https://github.com/vimsucks/wxwork-bot-go
	bot := wxworkbot.New("b2d57746-7146-44f2-8207-86cb0ca832be")
	markdown := wxworkbot.Markdown{
		Content: mdStr, // "# 测试",
	}
	err = bot.Send(markdown)
	if err != nil {
		log.Fatal(err)
	}
}
