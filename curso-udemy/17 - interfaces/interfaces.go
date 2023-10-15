package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("teste %f", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (c retangulo) area() float64 {
	return c.altura * c.largura
}

func (c circulo) area() float64 {
	return math.Pi * c.raio
}

func main() {

	r := retangulo{10, 15}

	c := circulo{10}

	escreverArea(r)
	escreverArea(c)

}
