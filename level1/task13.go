package main

import "fmt"

func main() {
	for {
		var x string
		fmt.Scan(&x)
		if x == "да" || x == "нет" || x == "чёрный" || x == "белый" {
			fmt.Println("Поражение")
			break
		} else {
			fmt.Println("Игра продолжается")
		}
	}
}
