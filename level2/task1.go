package main

import "fmt"

func GoOrNot(x string) {
	if x != "Go" {
		fmt.Println("Я знаю только Go!")
	} else {
		fmt.Println("Go!")
	}
}
