package main

import "errors"

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	i := 0
	var err error
	err = nil
	if n <= 0 {
		err = errors.New("hui")
	}
	ans := make([]int, 0, len(nums))
	for i < len(nums) && n > 0 {
		if nums[i] < limit {
			n--
			ans = append(ans, nums[i])
		}
		i++
	}

	return ans, err
}
