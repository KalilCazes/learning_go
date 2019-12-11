package main

import (
	"fmt"
)

func main() {

	veiculo := struct {
		info   []string
		modelo map[string]int
	}{
		info: []string{"porta", "pneu", "retrovisor", "airbag"},
		modelo: map[string]int{
			"Gol":   20000,
			"Palio": 25000,
			"Siena": 30000,
		},
	}

	fmt.Println(veiculo)
}
