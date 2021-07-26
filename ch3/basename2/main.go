// basename remove componentes de diretório e um .sufixo.
// exemplos: a => 1, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import "strings"

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 se "/" não for encontrada
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
