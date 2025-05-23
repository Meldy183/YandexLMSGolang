package main

import "fmt"

func PrintFlightRow(num, from, to string, dur float32, recep, gate int, isReg bool) {
	if !isReg {
		fmt.Printf("| %s | %s—%s | %d регистрация продолжается |\n", num, from, to, recep)
		return
	}
	fmt.Printf("| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |\n", num, from, to, gate, dur)
}
