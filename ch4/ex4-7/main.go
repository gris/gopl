// reverse inverte uma fatia de bytes in-place
package main

import (
	"fmt"
	"unicode/utf8"
)

// Copied from https://github.com/xingdl2007/gopl-solutions/blob/master/ch4/4.7/main.go
func reverse2(in []byte) {
	// first treat as non utf8-encoded data
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	// try to decode according to utf8, then fix error
	i := 0
	for i < len(in) {
		var tryTwo, tryThree, tryFour bool
		for {
			r, s := utf8.DecodeRune(in[i:])
			if r != utf8.RuneError {
				i += s
				break
			} else {
				// try two byte length, swap two bytes
				if !tryTwo {
					tryTwo = true
					in[i], in[i+1] = in[i+1], in[i]
					continue
				}

				// try three byte length, swap three bytes
				if !tryThree {
					// cancel tryTwo side effect
					in[i], in[i+1] = in[i+1], in[i]
					tryThree = true
					in[i], in[i+2] = in[i+2], in[i]
					continue
				}

				// try four byte length, swap four bytes
				if !tryFour {
					// cancel tryThree side effect
					in[i], in[i+2] = in[i+2], in[i]

					tryFour = true
					in[i], in[i+1], in[i+2], in[i+3] = in[i+3], in[i+2], in[i+1], in[i]
					continue
				}

				// should not be here
				panic("Should not be here!")
			}
		}
	}
}

func main() {
	s := []byte("hello 世界 world!")
	reverse2(s)
	fmt.Println(string(s))
}
