package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // deference p
	*p = 21         // set via dereference
	fmt.Println(i)  // check actual var

	p = &j         // point to j
	*p = *p / 37   // divide by dereferencing
	fmt.Println(j) // check actual var
}
