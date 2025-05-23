package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i, v := range array {
		if i <= 6 {
			fmt.Printf("%d я уже сделал: %s\n", i+1, v)
		} else {
			fmt.Printf("%d не успел сделать: %s\n", i+1, v)
		}
	}
}
