package main

import "fmt"

func main() {

	temp1 := incrementa()
	temp2 := incrementa()
	for x := 0; x < 6; x++ {
		fmt.Printf("\nTemp1: %d <-----> Temp2: %d\n", temp1(), temp2())
	}
}

func incrementa() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
