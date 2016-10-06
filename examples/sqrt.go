package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	diff, z := 1.0, x
	i := 0
	for next := x; diff > 0.00001; z = next {
		next = z - (z-x/z)/2
		diff = math.Abs(z - next)
		i++
	}
	fmt.Println("Iterations: ", i)
	return
}

func main() {
	fmt.Printf("%.6g\n", Sqrt(0.01))
}
