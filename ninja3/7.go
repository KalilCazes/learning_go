package main

import (
	"fmt"
)

func main() {

	number := -101

	if number > 0 && number%2 == 0 {
		fmt.Printf("%d %s\n", number, "even.")
	} else if number > 0 && number%2 != 0 {
		fmt.Printf("%d %s\n", number, "is odd.")
	} else {
		fmt.Printf("%d %s\n", number, "is negative.")
	}
}
