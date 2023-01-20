package persistence

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
	wxworkbot "github.com/vimsucks/wxwork-bot-go"
)

type Sites struct {
	List map[string]SiteStatus
}

type SiteStatus struct {
	Name string
	Code int
}

// Return value.
var robotRepository *Sites

func GetRobotRepository() *Sites {
	if robotRepository == nil {
		robotRepository = &Sites{}
	}
	return robotRepository
}

func (s *Sites) getWebSiteStatus() (*[]SiteStatus, *[]SiteStatus, error) {
	// http://c.biancheng.net/view/32.html
	healthySites := []SiteStatus{}
	failSites := []SiteStatus{}
	client := resty.New()
	// https://github.com/go-resty/resty/blob/master/redirect.go
	for url, status := range s.List {
		resCode := 0
		resp, err := client.R().
			Get(url)
		if err != nil {
			log.Println("error:", err)
		} else {
			resCode = resp.StatusCode()
		}
		if status.Code == resCode {
			healthySites = append(healthySites, status)
		} else {
			failSites = append(failSites, SiteStatus{status.Name, resCode})
		}
	}
	return &healthySites, &failSites, nil
}

func (s *Sites) ClearCheckResult() (*wxworkbot.Markdown, error) {
	ss := s
	ss.List = map[string]SiteStatus{
		"https://i.mazey.net/cdn/jquery-2.1.1.min.js": {"CDN jQuery", 200},
	}
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
		log.Println("error:", err)
	}
	mdStr := "Health Check Result:\n"
	for _, site := range *healthySites {
		mdStr += fmt.Sprintf("<font color=\"info\">%s OK</font>\n", site.Name)
	}
	lo.ForEach(*failSites, func(site SiteStatus, _ int) {
		// log.Println("ForEach SiteStatus:", site.Name)
		siteLink, _ := lo.FindKeyBy(ss.List, func(k string, v SiteStatus) bool {
			return v.Name == site.Name
		})
		// log.Println("siteLink:", siteLink)
		mdStr += fmt.Sprintf(
			"<font color=\"warning\">%s FAIL</font>\n"+
				"Error Code: %d\n"+
				"Link: [%s](%s)\n",
			site.Name,
			site.Code,
			siteLink,
			siteLink,
		)
	})
	mdStr += fmt.Sprintf("<font color=\"comment\">*%s%d*</font>", "Sum: ", len(*healthySites)+len(*failSites))
	persistenceGetAlias2dataRepository := GetAlias2dataRepository()
	data, err := persistenceGetAlias2dataRepository.Get("WECOM_ROBOT_CHECK")
	wxworkRobotKey := ""
	if err != nil {
		log.Println("error:", err)
		// Use ENV
		wxworkRobotKey = os.Getenv("WECOM_ROBOT_CHECK")
	} else {
		wxworkRobotKey = data.Data
	}
	// log.Println("Check data", data)
	// log.Println("Check WECOM_ROBOT_CHECK", data.Data)
	// https://github.com/vimsucks/wxwork-bot-go
	bot := wxworkbot.New(wxworkRobotKey)
	markdown := wxworkbot.Markdown{
		Content: mdStr,
	}
	err = bot.Send(markdown)
	if err != nil {
		log.Println("error:", err)
	}
	return &markdown, nil
}
