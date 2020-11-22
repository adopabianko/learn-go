package main

import (
	"fmt"
	"strings"
)

func IsPolindrome(str string) bool {
	var result string

	str = strings.ToLower(str)

	for _, v := range str {
		result = string(v) + result
	}

	if str != result {
		return false
	}

	return true
}

func main() {
	var str string

	str = "Katak"

	fmt.Println(str)

	if IsPolindrome(str) {
		fmt.Println("Is polindrome")
	} else {
		fmt.Println("Is not polindrome")
	}
}
