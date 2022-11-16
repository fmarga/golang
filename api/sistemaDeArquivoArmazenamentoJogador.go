package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SistemaDeArquivoArmazenamentoJogador struct {
	bancoDeDados *json.Encoder
	liga         Liga
}

func (s *SistemaDeArquivoArmazenamentoJogador) ObterLiga() Liga {
	sort.Slice(s.liga, func(i, j int) bool {
		return f.liga[i].Vitorias > f.liga[j].Vitorias
	})
	return s.liga
}

func (s *SistemaDeArquivoArmazenamentoJogador) PegarPontuacaoJogador(nome string) int {

	jogador := s.liga.Find(nome)
	if jogador != nil {
		return jogador.Vitorias
	}

	return 0
}

func (s *SistemaDeArquivoArmazenamentoJogador) SalvaVitoria(nome string) {
	jogador := s.liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		s.liga = append(s.liga, Jogador{nome, 1})
	}

	s.bancoDeDados.Encode(s.liga)
}

func NovoSistemaArquivoArmazenamentoJogador(arquivo *os.File) (*SistemaDeArquivoArmazenamentoJogador, err) {
	err := iniciaArquivoBDDojogador(arquivo)

	if err != nil {
		return nil, fmt.Errorf("problema carregando armazenamento de jogador do arquivo %s, %v", arquivo.Name(), err)
	}

	return &SistemaDeArquivoArmazenamentoJogador{
		bancoDeDados: json.NewEncoder(&fita{arquivo}),
		liga:         liga,
	}, nil,
}

func iniciaArquivoBDDoJogador(arquivo &os.File) error {
	arquivo.Seek(0, 0)

	info, err := arquivo.Stat()

	if err != nil {
		return nil, fmt.Errorf("problema carregando armazenamento de jogador do arquivo %s, %v", arquivo.Name(), err)
	}

	if info.Size() == 0 {
		arquivo.Write([]byte("[]"))
		arquivo.Seek(0, 0)
	}
	return nil
}