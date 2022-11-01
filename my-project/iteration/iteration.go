package iteration

import "testing"

// const quantidade = 5

func Repetir(caractere string, quantidade int) string {
	var repeticoes string
	for i := 0; i < quantidade; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}
