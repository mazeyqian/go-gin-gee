package persistence

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	ss.List = map[string]SiteStatus{}
	// ss.List = map[string]SiteStatus{
	// 	"https://i.mazey.net/cdn/jquery-2.1.1.min.js": {"CDN jQuery", 200},
	// }
	// List - begin
	// CDN
	ss.List["https://i.mazey.net/cdn/jquery-2.1.1.min.js"] = SiteStatus{"CDN/Net/Arc/jQuery", 200}
	ss.List["https://i.mazey.net/cdn/bootstrap-3.4.1/css/bootstrap.min.css"] = SiteStatus{"CDN/Net/Arc/Bootstrap CSS", 200}
	ss.List["https://mazey.cn/cdn/layer/layer.js"] = SiteStatus{"CDN/Cn/Layer", 200}
	ss.List["https://mazey.cn/cdn/jquery-2.1.1.min.js"] = SiteStatus{"CDN/Cn/jQuery", 200}
	// Website
	ss.List["https://blog.mazey.net/"] = SiteStatus{"Blog/Home", 200}
	ss.List[fmt.Sprintf("%s%d", "https://blog.mazey.net/?s=", time.Now().Unix())] = SiteStatus{"Blog/Search", 200}
	ss.List["https://www.zhibaifa.com/"] = SiteStatus{"White/Home", 200}
	ss.List["https://i.mazey.net/tool/markdown/"] = SiteStatus{"Tool/Arc/Markdown Converter", 200}
	// Api
	ss.List["https://mazey.cn/feperf/monitor/get/topic?userName=%E5%90%8E%E9%99%A4"] = SiteStatus{"Server/Monitor Topic", 200}
	ss.List["https://mazey.cn/server/nut/feeds?currentPage=1&pageSize=10&total=0&isPrivacy=1"] = SiteStatus{"Server/Nut Read Feeds", 200}
	ss.List["https://mazey.cn/t/k"] = SiteStatus{"Server/Tiny Redirect", 200}
	ss.List["https://mazey.cn/server/weather/new-daily?location=shanghai"] = SiteStatus{"Server/Weather Shanghai", 200}
	ss.List["https://mazey.cn/server/user/info"] = SiteStatus{"Server/Location", 200}
	ss.List["https://feperf.com/api/mazeychu/413fbc97-4bed-3228-9f07-9141ba07c9f3"] = SiteStatus{"Gee/Chu/NameId0920", 200}
	// List - end
	healthySites, failSites, err := ss.getWebSiteStatus()
	if err != nil {
		log.Println("error:", err)
	}
	sucessNames := []string{}
	lo.ForEach(*healthySites, func(site SiteStatus, _ int) {
		sucessNames = append(sucessNames, site.Name)
	})
	// Sort sucessNames
	sort.Strings(sucessNames)
	log.Println("sucessNames:", sucessNames)
	mdStr := "Health Check Result:\n"
	// for _, site := range *healthySites {
	// 	mdStr += fmt.Sprintf("<font color=\"info\">%s OK</font>\n", site.Name)
	// }
	lo.ForEach(sucessNames, func(name string, _ int) {
		mdStr += fmt.Sprintf("<font color=\"info\">%s OK</font>\n", name)
	})
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
	log.Println("Robot data:", data)
	wxworkRobotKey := ""
	if err != nil {
		log.Println("error:", err)
		// Use ENV
		wxworkRobotKey = os.Getenv("WECOM_ROBOT_CHECK")
		log.Println("Robot Getenv:", wxworkRobotKey)
	} else {
		wxworkRobotKey = data.Data
	}
	log.Println("Robot wxworkRobotKey:", wxworkRobotKey)
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
