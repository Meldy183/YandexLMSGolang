package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	x := len(a) >= 8
	y := len(b) >= 8
	switch {
	case x && y:
		fmt.Println("Оба пароля надёжные")
	case x:
		fmt.Println("Только первый пароль надёжный")
	case y:
		fmt.Println("Только второй пароль надёжный")
	default:
		fmt.Println("Оба пароля ненадёжные")
	}
}
