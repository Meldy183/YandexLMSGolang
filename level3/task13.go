package main

import (
	"maps"
	"slices"
)

func FindMaxKey(m map[int]int) int {
	keys := slices.Sorted(maps.Keys(m))
	return keys[len(keys)-1]
}
