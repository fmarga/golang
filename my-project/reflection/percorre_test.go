package reflection

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func TestPercorre(t *testing.T) {

	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"struct com um campo de string",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		},
		{
			"struct sem campo string",
			struct {
				Nome  string
				Idade int
			}{"Chris", 30},
			[]string{"Chris"},
		},
		{
			"campos aninhados",
			Pessoa{
				"Chris",
				Perfil{30, "Londres"},
			},
			[]string{"Chris", "Londres"},
		},
		{
			"ponteiros para coisas",
			&Pessoa{
				"Chris",
				Perfil{30, "Londres"},
			},
			[]string{"Chris", "Londres"},
		},
		{
			"slices",
			[]Perfil{
				{33, "Londres"},
				{34, "Reykjavík"},
			},
			[]string{"Londres", "Reykjavík"},
		},
		{
			"arrays",
			[2]Perfil{
				{33, "Londres"},
				{34, "Reykjavík"},
			},
			[]string{"Londres", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var resultado []string
			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas) {
				t.Errorf("resultado '%v', esperado '%v'", resultado, teste.ChamadasEsperadas)
			}
		})

		t.Run("com maps", func(t *testing.T) {
			mapA := map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			}

			var resultado []string
			percorre(mapA, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			verificaSeContem(t, resultado, "Bar")
			verificaSeContem(t, resultado, "Boz")
		})
	}

}

func verificaSeContem(t *testing.T, palheiro []string, agulha string) {
	contem := false

	for _, x := range palheiro {
		if x == agulha {
			contem = true
		}
	}

	if !contem {
		t.Errorf("esperava-se que '%+v' contivesse '%s', mas não continha", palheiro, agulha)
	}
}
