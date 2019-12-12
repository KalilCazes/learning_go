package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	p := pessoa{
		nome:      "Albus",
		sobrenome: "Dumbledore",
		idade:     150,
	}

	p.info()
}

func (p pessoa) info() {
	fmt.Printf("%s %s - %d anos\n", p.nome, p.sobrenome, p.idade)
}
