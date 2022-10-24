package main

import "fmt"

const x int = 40
const y = 10

func main() {
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T", y, y)
}
