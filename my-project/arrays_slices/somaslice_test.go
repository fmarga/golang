package arrays_slices

import "testing"

func TestSomaSlice(t *testing.T) {
	// t.Run("coleção de 5 números", func(t *testing.T) {
	// 	numeros := [5]int{1, 2, 3, 4, 5}

	// 	resultado := SomaSlice(numeros)
	// 	esperado := 15

	// 	if resultado != esperado {
	// 		t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
	// 	}
	// })

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4}

		resultado := SomaSlice(numeros)
		esperado := 10

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})
}
