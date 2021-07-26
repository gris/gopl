// Anagram informa se duas strings são anagramas uma da outra.
package main

import "fmt"

func anagram(s1, s2 string) bool {
	var countS1 = make(map[byte]int)
	var countS2 = make(map[byte]int)
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return false
	}

	for i := range s1 {
		countS1[s1[i]]++
		countS2[s2[i]]++
	}

	if len(countS1) != len(countS2) {
		return false
	}

	for i := range countS1 {
		if countS1[i] != countS2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(anagram("banana", "nabaka"))
}
