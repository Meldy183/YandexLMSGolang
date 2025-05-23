package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a == b && b == c {
		fmt.Println("Максимальное равенство")
	} else {
		fmt.Println("Не равны")
	}
}
