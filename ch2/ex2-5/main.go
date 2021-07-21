package popcount4

// PopCount devolve a população (número de bits definidos) de x
func PopCount(x uint64) int {
	var pop byte
	for {
		if x == 0 {
			break
		}
		pop += 1
		x = x & (x - 1)
	}
	return int(pop)
}
