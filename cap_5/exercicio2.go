package main

import "fmt"

func main() {
	x := 10
	y := 20
	z := 10

	a := x < y
	b := x <= y
	c := x == z
	d := y > z
	e := y >= x
	f := x != y

	fmt.Println(a, b, c, d, e, f)
}
