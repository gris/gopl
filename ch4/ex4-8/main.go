// Charcount conta caracteres Unicode
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int) // contagem dos caracteres Unicode

	invalid := 0 // contagem de caracteres UTF-8 inválidos

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // retorna runa, número de bytes, erro
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			counts["letter"]++

		case unicode.IsDigit(r):
			counts["digit"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsPunct(r):
			counts["punct"]++
		default:
			counts["other"]++
		}
	}
	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
