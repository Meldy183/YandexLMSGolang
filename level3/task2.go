package main

import (
	"unicode"
)

//	func main() {
//		CheckOnlyASCII("sf")
//	}
func CheckOnlyASCII(s string) bool {
	x := unicode.MaxASCII
	for _, v := range []rune(s) {
		if x < v {
			return false
		}
	}
	return true
}
