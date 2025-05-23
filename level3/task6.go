package main

import "slices"

func FindMinMaxInArray(array [10]int) (int, int) {
	return slices.Max(array[:]), slices.Min(array[:])
}
