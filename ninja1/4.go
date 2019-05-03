package main

import (
	"fmt"
)

type teste int

var x teste

func main() {

	fmt.Printf("tipo:%T\nvalor:%v\n", x, x)
	x = 42
	fmt.Printf("x:%d\n", x)
}
