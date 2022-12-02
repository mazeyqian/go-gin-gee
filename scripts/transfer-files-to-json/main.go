package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bitfield/script"
)

type GamecenterH5 struct {
	Viewurl     string `json:"viewurl"`
	Degradeurl  string `json:"degradeurl"`
	Manifesturl string `json:"manifesturl"`
	Bundleurl   string `json:"bundleurl"`
}

func main() {
	log.Println("- begin -")
	gh := &GamecenterH5{}
	domain := "//s1.hdslb.com/bfs/static/gameweb/gamecenter-h5/"
	projectPath := "./dist/gamecenter-h5"
	assignedProjects := "gamecenter-h5/gamecenter-h5"
	projects := []string{
		"placeholder",
	}
	regexStr := "^.+("
	for _, v := range projects {
		regexStr += fmt.Sprintf("%s|", v)
	}
	regexStr += fmt.Sprintf("%s).+$", assignedProjects)
	log.Println("regexStr:", regexStr)
	regex := regexp.MustCompile(regexStr)
	script.ListFiles(projectPath).MatchRegexp(regex).FilterLine(func(s string) string {
		fileName := filepath.Base(s)
		link := domain + fileName
		if strings.Contains(link, ".template.html") {
			gh.Viewurl = link
		}
		if strings.Contains(link, ".degrade.html") {
			gh.Degradeurl = link
		}
		if strings.Contains(link, ".client.") {
			gh.Manifesturl = link
		}
		if strings.Contains(link, ".server.") {
			gh.Bundleurl = link
		}
		return link
	}).Stdout()
	log.Println("gh:", *gh)
	file, _ := json.MarshalIndent(*gh, "", " ")
	_ = ioutil.WriteFile("./dist/gamecenter-h5/config---gamecenter-h5.json", file, 0644)
	log.Println("- end -")
}
