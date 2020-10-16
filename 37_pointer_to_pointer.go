package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a

	pptr = &ptr

	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Value available at **ptr = %d\n", **pptr)
}
