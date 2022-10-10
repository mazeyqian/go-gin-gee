package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type Sites struct {
	List map[string]int
}

func (s *Sites) getWebSiteStatus() ([]string, error) {
	okUrls := []string{}
	client := resty.New()
	for url, code := range s.List {
		log.Println("url:", url)
		log.Println("code expect:", code)
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
		if code == resCode {
			okUrls = append(okUrls, url)
		}
	}
	// resp, err := client.R().
	// 	Get(url)
	// if err != nil {
	// 	return []string{"0"}, err
	// }
	return okUrls, nil
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
	ss.List = make(map[string]int)
	ss.List["https://blog.mazey.net/"] = 200
	ss.List["https://tool.mazey.net/markdown/"] = 200
	okUrls, err := ss.getWebSiteStatus()
	if err != nil {
		log.Println("  Error      :", err)
	}
	log.Println("okUrls:", okUrls)
}
