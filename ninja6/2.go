package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 10, 50}
	fmt.Println(sum(values...))
	fmt.Println(anotherSum(values))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func anotherSum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
