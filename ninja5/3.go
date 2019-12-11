package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	c := caminhonete{
		veiculo: veiculo{
			portas: 2,
			cor:    "Verde",
		},
		tracaoNasQuatro: true,
	}

	s := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "Preto",
		},
		modeloLuxo: false,
	}
	fmt.Println(c)
	fmt.Printf("Caminhonete tem %d portas.\n", c.portas)
	fmt.Println(s)
	fmt.Printf("Sedan tem a cor %s.\n", s.cor)
}
