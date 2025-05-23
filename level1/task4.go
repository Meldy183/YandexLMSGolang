package main

import (
	"fmt"
	"math"
)

func main() {
	var q float64
	fmt.Scan(&q)
	sqrt := math.Sqrt(q)
	if q < 0 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%.3f", sqrt)
	}
}
