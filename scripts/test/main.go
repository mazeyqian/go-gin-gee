package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := int64(1686830930000)
	seconds := timestamp / 1000
	nanos := (timestamp % 1000) * 1000000
	t := time.Unix(seconds, nanos)
	hour := t.Format("15")
	fmt.Println(hour)
}
