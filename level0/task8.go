package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02"
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	var time1, time2 time.Time
	time1, _ = time.Parse(layout, str1)
	time2, _ = time.Parse(layout, str2)
	fmt.Printf("%d year ago", time1.Year()-time2.Year())
}
