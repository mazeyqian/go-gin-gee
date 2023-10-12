package main

import (
	"log"
	"net/url"
)

func main() {
	var testUrl0 string = "https://www.example.com:4433/path/to/somewhere?param1=1&param2=2"
	result, err := url.Parse(testUrl0)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Scheme:", result.Scheme)
	log.Println("Host:", result.Host)
	log.Println("Path:", result.Path)
	log.Println("RawQuery:", result.RawQuery)
}
