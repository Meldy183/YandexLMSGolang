package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	if str == "0" {
		fmt.Println("Стоит надеть куртку")
	} else {
		var x int
		fmt.Scan(&x)
		if str == "-" {
			x *= -1
		}
		switch {
		case x > 20:
			fmt.Println("Стоит надеть майку и шорты")

		case x >= 10:
			fmt.Println("Стоит надеть штаны и кофту")
		case x >= -5:
			fmt.Println("Стоит надеть куртку")
		default:
			fmt.Println("Стоит надеть зимнюю куртку")
		}
	}
}
