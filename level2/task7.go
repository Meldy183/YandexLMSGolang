package main

import (
	"strings"
)

func ratePassword(str string) string {
	ans := 0
	if hasMinimumLength(str, 8) {
		ans++
	}
	if hasUpper(str) {
		ans++
	}
	if hasLowerCase(str) {
		ans++
	}
	switch ans {
	case 0:
		return "Пароль недостаточно безопасен, придумайте новый"
	case 1:
		return "Слабый пароль"
	case 2:
		return "Средний пароль"
	case 3:
		return "Сложный пароль"
	}
	return str
}

func hasLowerCase(str string) bool {
	return strings.ToUpper(str) != str
}
