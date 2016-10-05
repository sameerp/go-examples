package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func add1(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 14))
}
