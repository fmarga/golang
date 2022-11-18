package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func criaArquivoTemporario(t *testing.T, dadoInicial string) (*os.File, func()) {
	t.Helper()

	arquivotmp, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("não foi possível escrever o arquivo tempor[ario %v", err)
	}

	arquivotmp.Write([]byte(dadoInicial))

	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}

func TestSistemaDeArquivoArmazenamentoJogador(t *testing.T) {
	// t.Run("/liga de um leitor", func(t *testing.T) {
	// 	bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
	// 		{"Nome": "Cleo", "Vitorias": 10},
	// 		{"Nome": "Chris", "Vitorias": 33}
	// 	]`)
	// 	defer limpaBancoDeDados()
	// 	armazenamento := SistemaDeArquivoArmazenamentoJogador{bancoDeDados}

	// 	recebido := armazenamento.ObterLiga()

	// 	esperado := []Jogador{
	// 		{"Cleo", 10},
	// 		{"Chris", 33},
	// 	}

	// 	verificaLiga(t, recebido, esperado)

	// 	recebido = armazenamento.ObterLiga()
	// 	verificaLiga(t, recebido, esperado)
	// })

	t.Run("retorna pontuação do jogador", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento, err := NovoSistemaArquivoArmazenamentoJogador(bancoDeDados)

		defineSemErro(t, err)

		recebido := armazenamento.ObterPontuacaoJogador("Chris")
		esperado := 33

		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitórias de um jogador existente", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento, err := NovoSistemaArquivoArmazenamentoJogador(bancoDeDados)

		defineSemErro(t, err)

		armazenamento.RegistrarVitoria("Chris")

		recebido := armazenamento.ObterPontuacaoJogador("Chris")
		esperado := 34
		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitórias de novos jogadores", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento, err := NovoSistemaArquivoArmazenamentoJogador(bancoDeDados)
		defineSemErro(t, err)

		armazenamento.RegistrarVitoria("Pepper")

		recebido := armazenamento.ObterPontuacaoJogador("Pepper")
		esperado := 1
		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("liga ordenada", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento, err := NovoSistemaArquivoArmazenamentoJogador(bancoDeDados)

		defineSemErro(t, err)

		recebido := armazenamento.ObterLiga()
		esperado := []Jogador{
			{"Chris", 33},
			{"Cleo", 10},
		}

		verificaLiga(t, recebido, esperado)

		recebido = armazenamento.ObterLiga()
		verificaLiga(t, recebido, esperado)
	})

	t.Run("trabalha com um arquivo vazio", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, "")
		defer limpaBancoDeDados()

		_, err := NovoSistemaArquivoArmazenamentoJogador(bancoDeDados)

		defineSemErro(t, err)
	})
}

func definePontuacaoIgual(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido %d, esperado %d", recebido, esperado)
	}
}

func defineSemErro(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("não esperava erro mas recebeu um, %v", err)
	}
}
