package main

import "fmt"

func main() {
	var a int = 10

	for a < 20 {
		if a == 15 {
			a = a + 1
			continue
		}

		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
