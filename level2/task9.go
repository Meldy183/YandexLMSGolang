package main

import "errors"

var DivisionByZeroError error

func Divide(a, b int) (float64, error) {
	if b == 0 {
		DivisionByZeroError = errors.New("division by zero is not allowed")
		return 0, DivisionByZeroError
	}
	return float64(a) / float64(b), nil
}
