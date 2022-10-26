package main

import "fmt"

func main() {
	x := 10
	y := 40

	switch {
	case x < y:
		fmt.Println("x is smaller than y")
	case x == y:
		fmt.Println("x is equal y")
	case x > y:
		fmt.Println("x is bigger than y")
	}
}
