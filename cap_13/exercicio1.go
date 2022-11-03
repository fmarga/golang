package main

import "fmt"

func retornaint() int {
	x := 3
	return x
}

func retornaintestring() (int, string) {
	x := 3
	y := "string"
	return x, y
}

func main() {
	fmt.Println(retornaint())
	fmt.Println(retornaintestring())
}
