package main

import "fmt"

type pessoa struct {
	nome                      string
	sobrenome                 string
	sabores_favoritos_sorvete []string
}

func main() {

	pessoa1 := pessoa{
		nome:                      "Maria",
		sobrenome:                 "Alves",
		sabores_favoritos_sorvete: []string{"flocos", "chocolate", "avelã"},
	}

	pessoa2 := pessoa{"Abigail", "Bittecourt", []string{"melancia", "abacaxi"}}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	for _, v := range pessoa1.sabores_favoritos_sorvete {
		fmt.Println(v)
	}
}
