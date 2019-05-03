package main

import (
	"fmt"
)

func main() {
	a := ("Alo" == "Alo")
	b := (10 != 10)
	c := (5 > 4)
	d := (5 < 4)
	e := (100 >= 100)
	f := (1 <= 10)
	fmt.Println(a, b, c, d, e, f)
}
