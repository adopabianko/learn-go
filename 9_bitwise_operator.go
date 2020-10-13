package main

import "fmt"

func main() {
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2
	fmt.Printf("Line 5 - Value of c is %d\n", c)
}
