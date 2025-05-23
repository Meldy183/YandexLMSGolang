package main

import "fmt"

func main() {
	var x int
	var d int
	fmt.Scan(&x, &d)
	ans := 0.0
	var v float64
	for range x {
		fmt.Scan(&v)
		ans += v
	}
	fmt.Println(float32(ans * (float64(100-d) / 100)))
}
