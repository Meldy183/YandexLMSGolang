package main

import "fmt"

func main() {
	var x float64
	var score int
	fmt.Scan(&score)
	for range score {
		fmt.Scan(&x)
		if x > 100 || x < 0 {
			fmt.Println("Неверный балл")
			continue
		}
		switch {
		case x >= 90:
			fmt.Println("5")
		case x >= 75:
			fmt.Println("4")
		case x >= 50:
			fmt.Println("3")
		case x >= 0:
			fmt.Println("2")
		}
	}
}
