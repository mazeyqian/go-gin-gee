package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chebyrash/promise"
)

func main() {
	p1 := promise.New(func(resolve func(int), reject func(error)) {
		factorial := findFactorial(20)
		resolve(factorial)
	})
	p2 := promise.New(func(resolve func(string), reject func(error)) {
		ip, err := fetchIP()
		if err != nil {
			reject(err)
			return
		}
		resolve(ip)
	})

	factorial, _ := p1.Await()
	fmt.Println(factorial)

	IP, _ := p2.Await()
	fmt.Println(IP)
}

func findFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * findFactorial(n-1)
}

func fetchIP() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		return "", err
	}

	type Response struct {
		Origin string `json:"origin"`
	}
	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Origin, err
}
