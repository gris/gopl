// rotate rotaciona uma fatia de n elementos in-place
package main

import "fmt"

func rotate(s []int) []int {
	return append(s[2:], s[:2]...)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = rotate(s)
	fmt.Println(s)
}
