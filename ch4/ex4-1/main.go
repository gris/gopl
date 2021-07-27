// Ex4-1 conta o número de bits diferentes em dois hashes SHA256
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(DiffBitsCount(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
}

func popcount(x uint8) byte {
	var pop byte
	for {
		if x == 0 {
			break
		}
		pop += 1
		x = x & (x - 1)
	}
	return pop
}

func DiffBitsCount(h1, h2 [32]uint8) int {
	var diffBits byte
	fmt.Println(h1)
	for i := range h1 {
		x := (h1[i] & h2[i]) | (^h1[i] & ^h2[i])
		diffBits += (8 - popcount(x))
	}
	return int(diffBits)
}
