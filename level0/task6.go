package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	fmt.Println(math.Round(max(x, y)))
}
