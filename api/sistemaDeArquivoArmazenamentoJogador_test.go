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
	t.Run("/liga de um leitor", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()
		armazenamento := SistemaDeArquivoArmazenamentoJogador{bancoDeDados}

		recebido := armazenamento.ObterLiga()

		esperado := []Jogador{
			{"Cleo", 10},
			{"Chris", 33},
		}

		verificaLiga(t, recebido, esperado)

		recebido = armazenamento.ObterLiga()
		verificaLiga(t, recebido, esperado)
	})

	t.Run("retorna pontuação do jogador", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoArmazenamentoJogador{bancoDeDados}
		recebido := armazenamento.PegarPontuacaoJogador("Chris")
		esperado := 33

		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitórias de um jogador existente", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoArmazenamentoJogador{bancoDeDados}
		armazenamento.SalvaVitoria("Chris")

		recebido := armazenamento.PegarPontuacaoJogador("Chris")
		esperado := 34
		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitórias de novos jogadores", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoArmazenamentoJogador{bancoDeDados}
		armazenamento.SalvaVitoria("Pepper")

		recebido := armazenamento.PegarPontuacaoJogador("Pepper")
		esperado := 1
		definePontuacaoIgual(t, recebido, esperado)
	})
}

func definePontuacaoIgual(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido %d, esperado %d", recebido, esperado)
	}
}
