package main

func DeleteLongKeys(m map[string]int) map[string]int {
	for k, _ := range m {
		if len(k) < 6 {
			delete(m, k)
		}
	}
	return m
}
