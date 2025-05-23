package main

import "math"

func SquareRoots(a, b, c float64) (float64, float64) {
	d := findDiscriminant(a, b, c)
	if d < 0 {
		return 0, 0
	}
	if d == 0 {
		ans := -b / (2 * a)
		return ans, ans
	}
	return (-b - math.Sqrt(d)) / (2 * a), (-b + math.Sqrt(d)) / (2 * a)
}

func findDiscriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}
