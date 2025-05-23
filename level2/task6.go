package main

import "strings"

func checkPassword(str string) bool {
	if !hasMinimumLength(str, 8) {
		return false
	}
	if !hasUpper(str) {
		return false
	}
	return true
}

func hasMinimumLength(str string, min int) bool {
	return len(str) >= min
}

func hasUpper(str string) bool {
	return strings.ToLower(str) != str
}
