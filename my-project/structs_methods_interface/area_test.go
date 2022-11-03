package structs_methods_interface

import "testing"

func TestArea(t *testing.T) {

	// Table driven test
	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retângulo", forma: Retangulo{largura: 12, altura: 6}, esperado: 72.0},
		{nome: "Círculo", forma: Circulo{raio: 10}, esperado: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{base: 12, altura: 6}, esperado: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.esperado {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
			}
		})
	}

	// teste antigo, separando por cada caso (forma)
	// verificaArea := func(t *testing.T, forma Forma, esperado float64) {
	// 	t.Helper()
	// 	resultado := forma.Area()

	// 	if resultado != esperado {
	// 		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	// 	}
	// }

	// t.Run("retângulos", func(t *testing.T) {
	// 	retangulo := Retangulo{2.0, 2.0}

	// 	verificaArea(t, retangulo, 4.0)

	// })

	// t.Run("círculos", func(t *testing.T) {
	// 	circulo := Circulo{10}

	// 	verificaArea(t, circulo, 314.1592653589793)
	// })
}
