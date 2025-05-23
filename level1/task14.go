package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var word string
	ans := 0
	fmt.Scan(&n, &word)
	word = strings.ToLower(word)
	cur := ""
	for range n {
		fmt.Scan(&cur)
		if strings.ToLower(cur) == word {
			ans++
		}
	}
	fmt.Println(ans)
}
