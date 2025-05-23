package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	newMap := make(map[string]string)
	for k, v := range m {
		newMap[v] = k
	}
	return newMap
}
