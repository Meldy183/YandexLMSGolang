package main

import "fmt"

func PrintComplexNumber(z complex64) {
	fmt.Printf("Действительная часть: %.2f. Мнимая часть: %.2f\n", real(z), imag(z))
}
