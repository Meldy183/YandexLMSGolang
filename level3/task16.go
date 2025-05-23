package main

func CountingSort(contacts []string) map[string]int {
	ans := make(map[string]int)
	for _, contact := range contacts {
		ans[contact]++
	}
	return ans
}
