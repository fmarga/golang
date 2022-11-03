package main

import "fmt"

func main() {
	defer fmt.Println("primeiro com defer")
	fmt.Println("depois sem defer")
}
