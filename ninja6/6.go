package main

import "fmt"

func main() {
	slice := []int{90, 80, 99, 1432}

	func(x ...int) {
		sum := 0
		for _, v := range x {
			sum += v
		}
		fmt.Println(sum)
	}(slice...)
}
