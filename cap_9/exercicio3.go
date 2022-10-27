package main

import "fmt"

func main() {
	slice := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	slice1 := slice[:3]
	slice2 := slice[4:]
	slice3 := slice[1:7]
	slice4 := slice[2:]
	slice5 := slice[2 : len(slice)-1]

	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
