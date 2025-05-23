package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02/15:04:05"
	var timer time.Time
	var str string
	fmt.Scanln(&str)
	timer, _ = time.Parse(layout, str)
	fmt.Printf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?", timer.Hour(), timer.Minute())
}
