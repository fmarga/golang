package arrays_slices

func Soma(arr []int) int {
	var resultado int
	for _, numero := range arr {
		resultado += numero
	}
	return resultado
}
