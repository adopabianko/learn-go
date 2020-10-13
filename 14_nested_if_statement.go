package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Printf("Value of a is 100 and b is 200\n")
		}
	}

	fmt.Printf("Exact value of a is : %d\n", a)
	fmt.Printf("Exact value of b is : %d\n", b)
}
