package main

import (
	"fmt"
)

func main() {

	for number := 1; number <= 20; number++ {

		if number%2 == 0 {
			fmt.Printf("%d %s\n", number, "is even.")
		} else {
			fmt.Printf("%d %s\n", number, "is odd.")
		}
	}
}
