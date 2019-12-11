package main

import "fmt"

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

	fmt.Printf("--------------> %s %s <--------------\n", p1.nome, p1.sobrenome)
	for _, v := range p1.sorvete {
		fmt.Printf("\t\t%s\n", v)
	}

	fmt.Printf("\n--------------> %s %s <--------------\n", p2.nome, p2.sobrenome)
	for _, v := range p2.sorvete {
		fmt.Printf("\t\t%s\n", v)
	}
}
