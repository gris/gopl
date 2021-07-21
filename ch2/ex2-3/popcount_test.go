package popcountloop_test

import (
	"testing"

	popcountloop "gopl.io/ch2/ex2-3"
	"gopl.io/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(25534)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountloop.PopCount(25534)
	}
}
