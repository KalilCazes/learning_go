package main

import (
	"fmt"
)

func main() {

	vector := [5]int{11, 22, 33, 44, 55}

	fmt.Printf("Array type:%T\n", vector)
	for _, v := range vector {
		fmt.Println(v)
	}
}
