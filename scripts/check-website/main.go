package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

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

}

func getWebSiteStatus(url string) (int, error) {
	client := resty.New()
	resp, err := client.R().
		Get(url)
	if err != nil {
		return 0, err
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:")
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :", resp)
	return resp.StatusCode(), err
}
