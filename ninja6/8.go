package main

import "fmt"

func main() {
	s := []int{10, 4, 6, 6}
	notaFinal := media(s...)

	fmt.Println(notaFinal(s...))
}

func media(x ...int) func(...int) int {
	return func(x ...int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total / len(x)
	}
}
