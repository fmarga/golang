package main

import "fmt"

func main() {
	for x := 3; x < 123; x++ {
		fmt.Printf("%d - %v\n", x, string(x))
	}
}
