package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	saldo Bitcoin
}

func (w *Wallet) Depositar(valor Bitcoin) {
	fmt.Printf("O endereço do saldo no Depositar é %v \n", &w.saldo)
	w.saldo += valor
}

func (w *Wallet) Saldo() Bitcoin {
	return w.saldo
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (w *Wallet) Retirar(valor Bitcoin) error {

	if valor > w.saldo {
		return ErroSaldoInsuficiente
	}

	w.saldo -= valor
	return nil
}
