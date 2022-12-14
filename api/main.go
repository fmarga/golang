package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problema abrindo %s %v", dbFileName, err)
	}

	armazenamento, err := NovoSistemaArquivoArmazenamentoJogador(db)

	if err != nil {
		log.Fatalf("problema criando sistema de arquivo de armazenamento, %v", err)
	}

	servidor := NovoServidorJogador(armazenamento)

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 '%v'", err)
	}
}
