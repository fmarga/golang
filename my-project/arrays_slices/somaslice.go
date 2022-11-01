package arrays_slices

func SomaSlice(slice []int) int {
	soma := 0

	for _, valor := range slice {
		soma += valor
	}
	return soma
}
