package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	x := strings.Count(text, "е")
	var ans string
	if x != 0 {
		ans = fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", x)
	} else {
		ans = "Текст готов к публикации!"
	}
	return ans
}
