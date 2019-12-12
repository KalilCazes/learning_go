package main

import (
	"fmt"
	"math"
)

type figura interface {
	area() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return c.raio * c.raio * math.Pi
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q := quadrado{
		lado: 9.1,
	}

	c := circulo{
		raio: 3.5,
	}

	fmt.Println(info(q))
	fmt.Println(info(c))
}
