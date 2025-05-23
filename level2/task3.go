package main

import "fmt"

func BuyFriesX(x string) int {
	var y int
	switch x {
	case "S":
		y = 49
	case "M":
		y = 89
	case "L":
		y = 99
	default:
		fmt.Println("Некорректный размер")
		return 0
	}
	return y
}

func BuyColaX(x string) int {
	var y int
	switch x {
	case "S":
		y = 99
	case "M":
		y = 119
	case "L":
		y = 139
	default:
		fmt.Println("Некорректный размер")
		return 0
	}
	return y
}
func BuyCola(x string) {
	num := BuyColaX(x)
	if num == 0 {
		return
	}
	printPrice(num, "Кола")
}
func BuyFries(x string) {
	num := BuyFriesX(x)
	if num == 0 {
		return
	}
	printPrice(num, "Картошка фри")
}
func printPrice(price int, name string) {
	fmt.Printf("%s будет стоить %d рублей\n", name, price)
}
