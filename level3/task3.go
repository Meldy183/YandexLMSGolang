package main

import (
	"strings"
)

//	func main() {
//		CheckOnlyASCII("sf")
//	}
func NumbersToLetters(input string) string {
	input = strings.ReplaceAll(input, "0", "ноль")
	input = strings.ReplaceAll(input, "1", "один")
	input = strings.ReplaceAll(input, "2", "два")
	input = strings.ReplaceAll(input, "3", "три")
	input = strings.ReplaceAll(input, "4", "четыре")
	input = strings.ReplaceAll(input, "5", "пять")
	input = strings.ReplaceAll(input, "6", "шесть")
	input = strings.ReplaceAll(input, "7", "семь")
	input = strings.ReplaceAll(input, "8", "восемь")
	input = strings.ReplaceAll(input, "9", "девять")
	input = strings.ReplaceAll(input, "*", "умножить на")
	input = strings.ReplaceAll(input, "+", "плюс")
	input = strings.ReplaceAll(input, "-", "минус")
	input = strings.ReplaceAll(input, "/", "разделить на")
	return input
}
