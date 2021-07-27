// Nonempty é um exemplo de um algoritmo in-place para fatias
package main

import "fmt"

// nonempty devolve uma fatia que armazena apenas as strings não vazias
// O array subjacente é modificado durante a chamada
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)            // `["one" "three" "three"]`
}

func nonempty2(strings []string) []string {
	out := strings[:0] // fatia de tamanho zero do original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
