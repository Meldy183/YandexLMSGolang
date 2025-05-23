package main

import "fmt"

func main() {
	word := ""
	fmt.Scan(&word)
	for _, letter := range word {
		if letter != 'а' && letter != 'о' {
			fmt.Print(string(letter))
		}
	}
}
