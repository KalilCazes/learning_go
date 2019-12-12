package main

import (
	"fmt"
)

func main() {

	palavra := ola("Mundo")
	fmt.Printf("%s %s\n", "Ola", palavra)
}

func ola(s string) string {
	return s
}
