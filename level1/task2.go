package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	if x > y {
		fmt.Println("Первое число больше второго")
	}
	if x < y {
		fmt.Println("Второе число больше первого")
	}
	if x == y {
		fmt.Println("Числа равны")
	}
}
