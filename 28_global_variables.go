package main

import "fmt"

var g int

func main() {
	var a, b int

	a = 10
	b = 20
	g = a + b

	fmt.Printf("value of a = %d, b = %d, g = %d \n", a, b, g)
}
