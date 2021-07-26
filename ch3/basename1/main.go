// basename remove componentes de diretório e um .sufixo.
// exemplos: a => 1, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

func basename(s string) string {
	// Descarta a última '/' e tudo que estiver antes
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserva tudo que estiver antes do último '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
