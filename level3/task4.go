package main

func FiveSteps(array [5]int) [5]int {
	x := make([]int, 0)
	for i := 0; i < len(array); i++ {
		x = append(x, array[5-i-1])
	}
	var ans [5]int
	copy(ans[:], x)
	return ans
}
