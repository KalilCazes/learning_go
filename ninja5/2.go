package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sorvete   []string
}

func main() {
	p1 := pessoa{
		nome:      "Severus",
		sobrenome: "Snape",
		sorvete:   []string{"Flocos", "Morango", "Abacaxi"},
	}

	p2 := pessoa{
		nome:      "Albus",
		sobrenome: "Dumbledore",
		sorvete:   []string{"Chocolate", "Napolitano", "Creme"},
	}
	valores := map[string]pessoa{
		"Snape":      p1,
		"Dumbledore": p2,
	}

	for i := range valores {
		fmt.Printf("\n%s %s", valores[i].nome, valores[i].sobrenome)
		for _, sabor := range valores[i].sorvete {
			fmt.Printf("\n--->\t%s", sabor)
		}
	}
}
