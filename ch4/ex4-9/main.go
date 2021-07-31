// Wordfreq conta frequências de palavras em um arquivo texto de entrada.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // contagem das palavras
	total := 0                     // número de palavras
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		w := in.Text()
		counts[w]++
		total++
	}
	fmt.Printf("category\tcount\n")
	for w, n := range counts {
		fmt.Printf("%q\t%g\n", w, float64(n)/float64(total))
	}
}
