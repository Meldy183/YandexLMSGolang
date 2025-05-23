package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanln(&x, &y, &z)
	if x+y/100 >= z {
		fmt.Print("Сегодня будет вкусный кофе!")
	} else {
		fmt.Println("Стоит подкопить")
	}
}
