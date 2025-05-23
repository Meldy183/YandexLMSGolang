package main

func SliceCopy(nums []int) []int {
	x := make([]int, len(nums))
	copy(x, nums)
	return x
}
