package main

import "fmt"

func main() {
	x := 10
	y := 20

	z := x < y

	if z {
		fmt.Println("x is smaller than y")
	} else {
		fmt.Println("x is bigger than y")
	}
}
