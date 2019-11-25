package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for unicode := 0; unicode < 3; unicode++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
