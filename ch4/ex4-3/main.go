// reverse inverte um ponteiro de array de ints in-place
package main

import "fmt"

func reverse(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	x := ([32]int{1, 2, 3, 4, 5, 6})
	reverse(&x)
	fmt.Println(x)
}
