package mocks

import (
	"bytes"
	"testing"
)

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

func (s SleeperSpy) Sleep() {
	s.Chamadas++
}

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	Contagem(buffer, sleeperSpy)

	resultado := buffer.String()
	esperado := `3
	2
	1
	Vai!`

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

	if sleeperSpy.Chamadas != 4 {
		t.Errorf("Não houve chamadas suficientes do sleeper, esperado 4, resultado %d", sleeperSpy.Chamadas)
	}
}
