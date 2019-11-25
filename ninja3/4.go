package main

import (
	"fmt"
)

func main() {

	year := 1996
	for {
		fmt.Println(year)
		if year == 2019 {
			break
		}
		year++
	}
}
