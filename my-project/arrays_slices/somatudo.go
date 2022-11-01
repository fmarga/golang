package arrays_slices

func SomaTudo(slices ...[]int) []int {
	var somas []int

	for _, numeros := range slices {
		somas = append(somas, Soma(numeros))
	}
	return somas
}
