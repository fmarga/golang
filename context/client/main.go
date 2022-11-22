package main

import (
	"context"
	"time"
	"net/http"
	"fmt"
	"os"
	"io"
	"log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(ctx)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer res.Body.Close()

	// pega o resultado da request e joga no stdout
	io.Copy(os.Stdout, res.Body)
}