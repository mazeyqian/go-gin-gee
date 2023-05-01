package main

import (
	"fmt"
	"strconv"
)

func ShanghaiToUTC(shanghaiTime string) (string, error) {
	shanghaiHour, err := strconv.Atoi(shanghaiTime[:2])
	if err != nil {
		return "", err
	}
	shanghaiMinute, err := strconv.Atoi(shanghaiTime[3:])
	if err != nil {
		return "", err
	}
	shanghaiTotalMinutes := shanghaiHour*60 + shanghaiMinute
	utcTotalMinutes := (shanghaiTotalMinutes - 480 + 1440) % 1440
	utcHour := utcTotalMinutes / 60
	utcMinute := utcTotalMinutes % 60
	utcTime := fmt.Sprintf("%02d:%02d", utcHour, utcMinute)
	return utcTime, nil
}

func main() {
	shanghaiTime := "10:05"
	utcTime, err := ShanghaiToUTC(shanghaiTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s in Shanghai is %s in UTC\n", shanghaiTime, utcTime)

	shanghaiTime = "04:01"
	utcTime, err = ShanghaiToUTC(shanghaiTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s in Shanghai is %s in UTC\n", shanghaiTime, utcTime)

	shanghaiTime = "02:59"
	utcTime, err = ShanghaiToUTC(shanghaiTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s in Shanghai is %s in UTC\n", shanghaiTime, utcTime)

	shanghaiTime = "00:05"
	utcTime, err = ShanghaiToUTC(shanghaiTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s in Shanghai is %s in UTC\n", shanghaiTime, utcTime)

	shanghaiTime = "00:65"
	utcTime, err = ShanghaiToUTC(shanghaiTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s in Shanghai is %s in UTC\n", shanghaiTime, utcTime)
}
