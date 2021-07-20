package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var input = []string{"lala", "ok", "lalaland", "tenta a sorte", "copihead"}

var input2 = append(input, input...)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		var s, sep string
		for i := 0; i < len(input2); i++ {
			s += sep + input2[i]
			sep = " "
		}
		fmt.Println(s)
		secs := time.Since(start).Seconds()
		fmt.Println(secs)
	}
}
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		s, sep := "", ""
		for _, arg := range input2 {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
		secs := time.Since(start).Seconds()
		fmt.Println(secs)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		fmt.Println(strings.Join(input2, " "))
		secs := time.Since(start).Seconds()
		fmt.Println(secs)
	}
}
