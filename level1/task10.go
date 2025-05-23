package main

import "fmt"

func main() {
	var x string
	for range 5 {
		fmt.Scan(&x)
		if x == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
