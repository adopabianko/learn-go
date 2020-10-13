package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("Line 1 - Condition is true")
	}

	if a || b {
		fmt.Println("Line 2 - Condition is true")
	}

	// Lets change the value of a and b
	a = false
	b = true
	if a && b {
		fmt.Println("Line 3 - Condition is true")
	} else {
		fmt.Println("Line 3 - Condition is false")
	}

	if !(a && b) {
		fmt.Println("Line 4 - Condition is true")
	}
}
