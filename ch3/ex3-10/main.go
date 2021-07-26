// comma insere vírgulas em uma string que representa um inteiro decimal
// não negativo
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
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
	fmt.Println(comma("12324344545"))
}
