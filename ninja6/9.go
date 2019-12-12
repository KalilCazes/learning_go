package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(imparesSomados(soma, slice...))
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func imparesSomados(f func(x ...int) int, y ...int) int {
	var s []int
	for _, v := range y {
		if v%2 != 0 {
			s = append(s, v)
		}
	}
	impares := soma(s...)
	return impares
}
