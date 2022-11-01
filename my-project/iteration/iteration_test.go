package iteration

import "testing"
import "fmt"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("a", 5)
	fmt.Println(repeticoes)
	// Output: aaaaa
}
