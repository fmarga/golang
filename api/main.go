package main

import (
	"log"
	"net/http"
)

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return a.armazenamento[nome]
}

func main() {
	servidor := &ServidorJogador{NovoArmazenamentoJogadorEmMemoria()}

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 '%v'", err)
	}
}
