package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		resto := i % 4
		fmt.Println(resto)
	}
}
