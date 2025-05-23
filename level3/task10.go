package main

func Join(nums1, nums2 []int) []int {
	var x = make([]int, len(nums1)+len(nums2))
	nums1 = append(nums1, nums2...)
	copy(x, nums1)
	return x
}
