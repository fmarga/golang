package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes       map[string]int
	registroVitorias []string
	liga             []Jogador
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func (e *EsbocoArmazenamentoJogador) RegistrarVitoria(nome string) {
	e.registroVitorias = append(e.registroVitorias, nome)
}

func (e *EsbocoArmazenamentoJogador) ObterLiga() Liga {
	return e.liga
}

// ------ TESTE OBTER JOGADORES
func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
		nil,
		nil,
	}

	servidor := NovoServidorJogador(&armazenamento)

	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoRequisicao(t, resposta.Body.String(), "20")
	})

	t.Run("retornar resultado de Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoRequisicao(t, resposta.Body.String(), "10")
	})

	t.Run("retorna 404 para jogador não encontrado", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Jorge")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusNotFound)
	})
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisição é inválido, obtive '%s', esperava '%s'", recebido, esperado)
	}
}

func verificarRespostaCodigoStatus(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("não recebeu código de status HTTP esperado, recebeu '%d', esperava '%d'", recebido, esperado)
	}
}

// ------ TESTE ARMAZENAMENTO VITORIAS
func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil,
		nil,
	}
	servidor := NovoServidorJogador(&armazenamento)

	t.Run("registra vitórias nas chamadas ao método HTTP POST", func(t *testing.T) {
		jogador := "Maria"

		requisicao := novaRequisicaoRegistrarVitoriaPost(jogador)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusAccepted)

		if len(armazenamento.registroVitorias) != 1 {
			t.Errorf("verifiquei %d chamadas ao RegistrarVitoria, esperava %d", len(armazenamento.registroVitorias), 1)
		}

		if armazenamento.registroVitorias[0] != jogador {
			t.Errorf("não registrou o jogador corretamente, recebi '%s', esperava '%s'", armazenamento.registroVitorias[0], jogador)
		}

	})
}

func novaRequisicaoRegistrarVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

// ------ TESTE LIGA

func TestLiga(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{}
	servidor := NovoServidorJogador(&armazenamento)

	t.Run("retorna 200 em /liga", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		var obtido []Jogador

		err := json.NewDecoder(resposta.Body).Decode(&obtido)

		if err != nil {
			t.Fatalf("não foi possível fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", resposta.Body, err)
		}

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
	})

	t.Run("retorna a tabela da liga como JSON", func(t *testing.T) {
		ligaEsperada := []Jogador{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		armazenamento := EsbocoArmazenamentoJogador{nil, nil, ligaEsperada}
		servidor := NovoServidorJogador(&armazenamento)

		requisicao := novaRequisicaoDeLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		obtido := obterLigaDaResposta(t, resposta.Body)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificaLiga(t, obtido, ligaEsperada)
		verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)
	})
}

func novaRequisicaoDeLiga() *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
	return requisicao
}

func obterLigaDaResposta(t *testing.T, body io.Reader) (liga []Jogador) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&liga)
	if err != nil {
		t.Fatalf("não foi possível fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", body, err)
	}
	return
}

func verificaLiga(t *testing.T, obtido, esperado []Jogador) {
	t.Helper()
	if !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("obtido %v, esperado %v", obtido, esperado)
	}
}

const tipoDoConteudoJSON = "application/json"

func verificaTipoDoConteudo(t *testing.T, resposta *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	if resposta.Result().Header.Get("content-type") != esperado {
		t.Errorf("resposta não tinha o tipo de conteúdo de application/json, obtido '%v'", resposta.Result().Header)
	}
}
