package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(10, 5))
	fmt.Println(odd(sum(10, 5)))
}

func sum(a, b int) int {
	return a + b
}

func odd(x int) (int, string) {
	if x%2 == 0 {
		return x, "is odd"
	}
	return x, "is not odd"
}
