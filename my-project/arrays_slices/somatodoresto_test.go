package arrays_slices

import "testing"
import "reflect"

func TestSomaTodoResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		resultado := SomaTodoResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}

		verificaSomas(t, resultado, esperado)
	})
}
