package popcount3

// PopCount devolve a população (número de bits definidos) de x
func PopCount(x uint64) int {
	var pop byte
	for i := 0; i < 64; i++ {
		if x%2 == 1 {
			pop += 1
		}
		x >>= 1
	}
	return int(pop)
}
