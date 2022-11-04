package pointers_errors

import "testing"

func TestWallet(t *testing.T) {

	confirmaSaldo := func(t *testing.T, wallet Wallet, esperado Bitcoin) {
		t.Helper()
		resultado := wallet.Saldo()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	confirmaErro := func(t *testing.T, resultado error, esperado error) {
		t.Helper()

		if resultado == nil {
			t.Fatal("Esperava um erro mas nada ocorreu")
		}

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	confirmaErroInexistente := func(t *testing.T, resultado error) {
		t.Helper()
		if resultado != nil {
			t.Fatal("erro inesperado recebido")
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Depositar(Bitcoin(10))

		confirmaSaldo(t, wallet, Bitcoin(10))
		// fmt.Printf("O endereço do saldo do teste é %v \n", &wallet.saldo)
		// // vê o espaço da memoria que ocupa essa variavel
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		wallet := Wallet{saldo: Bitcoin(20)}

		erro := wallet.Retirar(Bitcoin(10))

		confirmaSaldo(t, wallet, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		wallet := Wallet{saldo: saldoInicial}
		erro := wallet.Retirar(Bitcoin(100))

		confirmaSaldo(t, wallet, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})

}
