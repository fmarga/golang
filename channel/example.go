package main

import "fmt"

func main() {
	messages := make(chan string)
	// cria um canal de string

	go func() {
		messages <- "ping"
	}()
	// envia a mensagem "ping" para o canal

	msg := <-messages
	// a variavel msg recebe o valor do canal messages

	fmt.Println(msg)
}