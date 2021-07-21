package popcountloop

// pc[i] é a população de i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount devolve a população (número de bits definidos) de x
func PopCount(x uint64) int {
	var pop byte
	for i := 0; i < 8; i++ {
		pop += pc[byte(x>>(i*8))]
	}
	return int(pop)
}
