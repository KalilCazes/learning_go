package main

import (
	"fmt"
)

const w int = 10
const x float64 = 10.5
const y bool = true
const z string = "ola"
const a = 1
const b = 1.5
const c = false
const d = "tchau"

func main() {
	fmt.Printf("%v - %v - %v - %v\n", w, x, y, z)
	fmt.Println("#######################")
	fmt.Printf("%v - %v - %v - %v\n", a, b, c, d)
}
