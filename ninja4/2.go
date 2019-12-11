package main

import "fmt"

func main() {

	slice := []int{}

	for i := 2; i < 21; i = i + 2 {
		slice = append(slice, i)
	}

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Printf("Slice:%T\n", slice)
}
