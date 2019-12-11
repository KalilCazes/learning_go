package main

import (
	"fmt"
)

func main() {

	slice := make([]string, 26, 26)
	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia",
		"Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná",
		"Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
		"Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Printf("Length:%d\nCapacity:%d\n", len(slice), cap(slice))

	fmt.Println(slice)
}
