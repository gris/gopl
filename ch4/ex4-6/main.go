// Ex4-6 elimina sequências de espaço Unicode adjacentes em uma fatia de bytes utf-8.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func eliminateDuplicateSpaces(s []byte) []byte {
	out := s[:0]
	var last rune

	for i := 0; i < len(s); {
		c, length := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(c) {
			out = append(out, s[i:i+length]...)
		} else if unicode.IsSpace(c) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}
		last = c
		i += length
	}
	return out
}

func main() {
	s := []byte(" \n\t\n⌘ okdsdfsdsds \n\t")
	fmt.Println(string(eliminateDuplicateSpaces(s)))
}
