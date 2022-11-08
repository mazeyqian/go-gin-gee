package schedules

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	wxworkbot "github.com/vimsucks/wxwork-bot-go"
)

type Sites struct {
	List map[string]SiteStatus
}

type SiteStatus struct {
	Name string
	Code int
}

func (s *Sites) getWebSiteStatus() ([]string, []string, error) {
	// http://c.biancheng.net/view/32.html
	healthySites := []string{}
	failSites := []string{}
	client := resty.New()
	// https://github.com/go-resty/resty/blob/master/redirect.go
	// client.SetRedirectPolicy(resty.NoRedirectPolicy())
	for url, status := range s.List {
		// log.Println("url:", url)
		// log.Println("code expect:", status.Code)
		resCode := 0
		resp, err := client.R().
			Get(url)
		if err != nil {
			// return 0, err
			log.Println("error:", err)
		} else {
			resCode = resp.StatusCode()
			// log.Println("code get:", resp.StatusCode())
		}
		if status.Code == resCode {
			healthySites = append(healthySites, status.Name)
		} else {
			failSites = append(failSites, status.Name)
		}
	}
	// resp, err := client.R().
	// 	Get(url)
	// if err != nil {
	// 	return []string{"0"}, err
	// }
	return healthySites, failSites, nil
}

func Check() {
	ss := &Sites{}
	ss.List = make(map[string]SiteStatus)
	ss.List["https://blog.mazey.net/"] = SiteStatus{"Blog Home", 200}
	ss.List[fmt.Sprintf("%s%d", "https://blog.mazey.net/?s=", time.Now().Unix())] = SiteStatus{"Blog Search", 200}
	ss.List["https://feperf.com/api/mazeychu/413fbc97-4bed-3228-9f07-9141ba07c9f3"] = SiteStatus{"Gee Gin NameId0920", 200}
	ss.List["https://i.mazey.net/tool/markdown/"] = SiteStatus{"Markdown Converter", 200}
	ss.List["https://mazey.cn/feperf/monitor/get/topic?userName=%E5%90%8E%E9%99%A4"] = SiteStatus{"FE Perf Monitor Topic", 200}
	ss.List["https://mazey.cn/server/nut/feeds?currentPage=1&pageSize=10&total=0&isPrivacy=1"] = SiteStatus{"Nut Read Feeds", 200}
	ss.List["https://mazey.cn/t/k"] = SiteStatus{"Tiny Redirect", 200}
	ss.List["https://www.zhibaifa.com/"] = SiteStatus{"Treat White Home", 200}
	ss.List["https://mazey.cn/server/weather/new-daily?location=shanghai"] = SiteStatus{"Weather Shanghai", 200}
	ss.List["https://mazey.cn/server/user/info"] = SiteStatus{"Location", 200}
	healthySites, failSites, err := ss.getWebSiteStatus()
	if err != nil {
		log.Println("  Error      :", err)
	}
	// log.Println("Healthy Sites:", healthySites)
	mdStr := "Health Check Result:\n"
	for _, siteName := range healthySites {
		mdStr += fmt.Sprintf("<font color=\"info\">%s OK</font>\n", siteName)
	}
	for _, siteName := range failSites {
		mdStr += fmt.Sprintf("<font color=\"warning\">%s FAIL</font>\n", siteName)
	}
	mdStr += fmt.Sprintf("<font color=\"comment\">*%s%d*</font>", "Sum: ", len(healthySites)+len(failSites))
	// log.Println(mdStr)
	s := persistence.GetAlias2dataRepository()
	data, err := s.Get("WECOM_ROBOT_CHECK")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Check data", data)
	log.Println("Check WECOM_ROBOT_CHECK", data.Data)
	// https://github.com/vimsucks/wxwork-bot-go
	bot := wxworkbot.New(data.Data) // "b2d57746-7146-44f2-8207-86cb0ca832be")
	markdown := wxworkbot.Markdown{
		Content: mdStr, // "# 测试",
	}
	err = bot.Send(markdown)
	if err != nil {
		log.Fatal(err)
	}
}

func RunCheck() {
	// https://github.com/go-co-op/gocron
	// s := gocron.NewScheduler(time.UTC)
	// https://pkg.go.dev/time#Location
	shanghai, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(shanghai)
	// s.Every(10).Seconds().Do(Check)
	s.Every(1).Day().At("10:00").Do(Check)
	s.StartAsync()
	// Check()
}
