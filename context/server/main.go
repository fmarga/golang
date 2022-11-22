package main

import(
	"net/http"
	"fmt"
	"time"
	"log"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// usa o context da propria request
	ctx := r.Context()
	log.Println("Iniciou minha request")
	defer log.Println("Finalizou minha request")

	select {
	// se chegar a 5s a requisição é feita
	case <-time.After(5 * time.Second):
		log.Println("requisição processada com sucesso")
		fmt.Fprintln(w, "requisição processada com sucesso")
		w.Write([]byte("requisição processada com sucesso"))
	// se algo acontecer no meio, a requisição é cancelada
	case <-ctx.Done():
		http.Error(w, "request cancelada", http.StatusRequestTimeout)
	}
}