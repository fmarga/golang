package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type SistemaDeArquivoArmazenamentoJogador struct {
	bancoDeDados *json.Encoder
	liga         Liga
}

func (s *SistemaDeArquivoArmazenamentoJogador) ObterLiga() Liga {
	sort.Slice(s.liga, func(i, j int) bool {
		return s.liga[i].Vitorias > s.liga[j].Vitorias
	})
	return s.liga
}

func (s *SistemaDeArquivoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {

	jogador := s.liga.Find(nome)
	if jogador != nil {
		return jogador.Vitorias
	}

	return 0
}

func (s *SistemaDeArquivoArmazenamentoJogador) RegistrarVitoria(nome string) {
	jogador := s.liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		s.liga = append(s.liga, Jogador{nome, 1})
	}

	s.bancoDeDados.Encode(s.liga)
}

func NovoSistemaArquivoArmazenamentoJogador(arquivo *os.File) (*SistemaDeArquivoArmazenamentoJogador, error) {
	err := iniciaArquivoBDDoJogador(arquivo)

	liga, err := NovaLiga(arquivo)

	if err != nil {
		return nil, fmt.Errorf("problema carregando armazenamento de jogador do arquivo %s, %v", arquivo.Name(), err)
	}

	return &SistemaDeArquivoArmazenamentoJogador{
		bancoDeDados: json.NewEncoder(&fita{arquivo}),
		liga:         liga,
	}, nil
}

func iniciaArquivoBDDoJogador(arquivo *os.File) error {
	arquivo.Seek(0, 0)

	info, err := arquivo.Stat()

	if err != nil {
		return fmt.Errorf("problema carregando armazenamento de jogador do arquivo %s, %v", arquivo.Name(), err)
	}

	if info.Size() == 0 {
		arquivo.Write([]byte("[]"))
		arquivo.Seek(0, 0)
	}
	return nil
}
