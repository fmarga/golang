package main

import "fmt"

func main() {
	estados := make([]string, 26, 26)
	fmt.Println(len(estados), cap(estados))
	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}
