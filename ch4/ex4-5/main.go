// Ex4-5 elimina duplicatas adjacentes em uma fatia de strings.
package main

import "fmt"

func eliminateDuplicates(s []string) []string {
	i := 0
	for {
		if i == 0 {
			i += 1
			continue
		}
		if i >= len(s) {
			return s
		}
		if s[i] == s[i-1] {
			s = remove(s, i)
			continue
		} else {
			i += 1
		}
	}
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	s := []string{"ok", "s", "s", "s", "haha", "haha", "s", "legal", "ok", "ok"}
	fmt.Println(eliminateDuplicates(s))
}
