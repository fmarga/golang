package main

import "fmt"

func main() {
	x := 1994
	y := 2022
	for {
		if x > y {
			break
		}
		fmt.Println(x)
		x++
	}
}
