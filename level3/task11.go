package main

func Mix(nums []int) []int {
	var result = make([]int, 0, len(nums))
	for i, elem := range nums[:len(nums)/2] {
		result = append(result, elem)
		result = append(result, nums[i+len(nums)/2])
	}
	return result
}
