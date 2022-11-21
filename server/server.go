package main

import (
	"fmt"
	"net/http"
	"log"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Testing a new server\n")
}

func main() {
	http.HandleFunc("/testing", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}