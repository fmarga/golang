package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5}
	sidois := []int{2, 4, 6, 8, 10}

	fmt.Println(umafuncao((si...)))
	fmt.Println(outrafuncao((sidois)))
}

func umafuncao(x ...int) int {
	total := 0
	for _, i := range x {
		total += i
	}
	return total
}

func outrafuncao(x []int) int {
	total := 0
	for _, i := range x {
		total += i
	}
	return total
}
