package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	if str == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}
