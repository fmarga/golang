package structs_methods_interface

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.altura + retangulo.largura)
}
