package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa) {
	p.idade += 200
}

func main() {

	p1 := pessoa{
		nome:  "Alvus",
		idade: 150,
	}
	fmt.Println(p1)
	mudeMe(&p1)
	fmt.Println(p1)
}
