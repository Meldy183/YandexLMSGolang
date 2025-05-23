package main

import "fmt"

func main() {
	var x, y string
	fmt.Scan(&x, &y)
	switch x {
	case "камень":
		x = "s"
	case "ножницы":
		x = "n"
	case "бумага":
		x = "p"
	}
	switch y {
	case "камень":
		y = "s"
	case "ножницы":
		y = "n"
	case "бумага":
		y = "p"
	}
	if x == y {
		fmt.Println("Ничья")
		return
	}
	if x == "s" && y == "n" || x == "n" && y == "p" || x == "p" && y == "s" {
		fmt.Println("Первый игрок победил")
	} else {
		fmt.Println("Второй игрок победил")
	}
}
