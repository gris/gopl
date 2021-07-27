// Concatena um inteiro a uma fatia de inteiros
package main

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Há espaço para crescer. Estende a fatia
		z = x[:zlen]
	} else {
		// Não há espaço suficiente. Aloca um novo array
		// Cresce para o dobro, para complexidade linear amortizada
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Uma função embutida; veja o texto
	}
	z[len(x)] = y
	return z
}
