package main

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func NovoArmazenamentoJogadorEmMemoria() *ArmazenamentoJogadorEmMemoria {
	return &ArmazenamentoJogadorEmMemoria{map[string]int{}}
}

func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return a.armazenamento[nome]
}

func (a *ArmazenamentoJogadorEmMemoria) ObterLiga() Liga {
	var liga []Jogador
	for nome, vitorias := range a.armazenamento {
		liga = append(liga, Jogador{nome, vitorias})
	}
	return liga
}
