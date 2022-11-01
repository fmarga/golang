package arrays_slices

import "testing"

func TestSoma(t *testing.T) {
	numeros := []int{1, 2, 3, 4, 5}
	resultado := Soma(numeros)
	esperado := 15

	if esperado != resultado {
		t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
	}
}
