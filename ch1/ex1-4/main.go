// Dup2 exibe a contagem, o texto das linhas que aparecem mais de uma vez na entrada e o nome dos arquivos.
// Ele lê de stdin ou de uma lista de arquivos nomeados.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, strings.Join(fileNames[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNames[input.Text()] = append(fileNames[input.Text()], f.Name())
	}
	// NOTA: ignorando erros em potencial de input.Err()
}
