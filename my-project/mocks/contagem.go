package mocks

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}

const (
	inicioContagem = 3
	ultimaPalavra  = "Vai!"
)

type SleeperPadrao struct{}

func (s SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}

	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}
