package main

import (
	"fmt"
)

type teste int

var x teste
var y int

func main() {

	fmt.Printf("tipo:%T\nvalor:%v\n", x, x)
	x = 42
	fmt.Printf("x:%d\n", x)
	y = int(x)
	fmt.Printf("valor de y:%v\ntipo de y:%T\n", y, y)
}
