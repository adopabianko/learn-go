package main

import "fmt"

func main() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("Line 1 - = Operator Example, Value of c = %d\n", c)

	c += a
	fmt.Printf("Line 2 - += Operator Example, Value of c = %d\n", c)

	c -= a
	fmt.Printf("Line 3 - -= Operator Example, Value of c = %d\n", c)

	c *= a
	fmt.Printf("Line 4 - *= Operator Example, Value of c = %d\n", c)

	c /= a
	fmt.Printf("Line 5 - /= Operator Example, Value of c = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("Line 5 - <<= Operator Example, Value of c = %d\n", c)

	c &= 2
	fmt.Printf("Line 5 - &= Operator Example, Value of c = %d\n", c)

	c ^= 2
	fmt.Printf("Line 5 - ^= Operator Example, Value of c = %d\n", c)

	c |= 2
	fmt.Printf("Line 5 - |= Operator Example, Value of c = %d\n", c)
}
