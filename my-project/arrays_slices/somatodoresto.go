package arrays_slices

func SomaTodoResto(slices ...[]int) []int {
	var somas []int

	for _, numeros := range slices {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}
	return somas
}
