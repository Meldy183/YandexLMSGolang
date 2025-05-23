package main

import (
	"fmt"
)

func main() {
	var q int
	fmt.Scan(&q)
	if q == 0 {
		fmt.Println("Число 0")
		return
	}
	if q > -10 && q < 10 {
		fmt.Println("Число однозначное")
		return
	}
	if q%2 == 0 {
		fmt.Println("Число чётное")
		return
	}
	if q > 0 {
		fmt.Println("Число положительное")
		return
	}
	fmt.Println("Число красивое")
}
