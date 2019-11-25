package main

import (
	"fmt"
)

func main() {

	number := -10
	switch {
	case number < 0:
		fmt.Printf("%d %s\n", number, "is negative")

	case number%2 == 0:
		fmt.Printf("%d %s\n", number, "is even")

	case number%2 != 0:
		fmt.Printf("%d %s\n", number, "is odd")
	}
}
