package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	* &a indicates pointer to a ie. address of variable a and
	* &b indicates pointer to b ie. address of variable b.
	 */
	swap(&a, &b)

	fmt.Printf("After swap, value of a: %d\n", a)
	fmt.Printf("After swap, value of b: %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
