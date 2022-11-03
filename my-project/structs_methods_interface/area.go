package structs_methods_interface

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	largura, altura float64
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return c.raio * c.raio * math.Pi
}

type Triangulo struct {
	base, altura float64
}

func (t Triangulo) Area() float64 {
	return (t.base * t.altura) / 2
}

// não é mais necessário usar essa funcao
// func Area(retangulo Retangulo) float64 {

// 	area := retangulo.altura * retangulo.largura

// 	return area
// }
