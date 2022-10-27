package main

import "fmt"

func main() {

	slice := [][]string{
		[]string{
			"maria",
			"vargas",
			"ler"
		},
		[]string{
			"joana",
			"vasconcellos",
			"dormir"
		},
		[]string{
			"leila",
			"camargo",
			"dançar"
		},
	}
	for _, v := range slice {
		fmt.Println(v[0])
		for _, i := range v {
			fmt.Println("\t", v)
		}
	}
}
