package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Offset returns the starting number of result for pagination
func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

// SortOrder returns the string for sorting and orderin data
func SortOrder(table, sort, order string) string {
	return table + "." + ToSnakeCase(sort) + " " + ToSnakeCase(order)
}

// ToSnakeCase changes string to database table
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func ConvertStringToMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ConvertShanghaiToUTC(shanghaiTime string) (string, error) {
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
