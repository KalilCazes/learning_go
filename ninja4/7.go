package main

import (
	"fmt"
)

func main() {

	matrix := [][]string{
		[]string{"Tony", "Hawk", "Skate"},
		[]string{"Bob", "Dylan", "Guitar"},
		[]string{"Freddie", "Mercury", "Sing"},
	}

	for _, v := range matrix {
		fmt.Println()
		for _, i := range v {

			fmt.Printf("%s\t\t\t", i)
		}

	}
}
