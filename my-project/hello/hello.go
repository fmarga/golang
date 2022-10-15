package main

import "fmt"

// this means this function returns a string
func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}
