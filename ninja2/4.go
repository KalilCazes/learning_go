package main

import (
	"fmt"
)

var x int = 10
var y int

func main() {
	fmt.Printf("decimal: %v\nbinario: %b\nhexadecimal: %#x\n", x, x, x)
	y = x << 1
	fmt.Println("-------------------------------")
	fmt.Printf("decimal: %v\nbinario: %b\nhexadecimal: %#x\n", y, y, y)

}
