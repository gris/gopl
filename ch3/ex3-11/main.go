// comma insere vírgulas em uma string que representa um inteiro decimal
// não negativo
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	signal := len(s) > 0 && s[0] == '+' || s[0] == '-'
	if signal {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		n := len(s[:dot])
		mod := n % 3
		for i, v := range s[:dot] {
			if i != 0 && i%3 == mod {
				buf.WriteRune(',')
			}
			buf.WriteRune(v)
		}
		buf.WriteRune('.')
		s = s[dot+1:]
	}
	n := len(s)
	mod := n % 3
	for i, v := range s {
		if i != 0 && i%3 == mod {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("+1232434224.5425"))
}
