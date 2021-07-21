package popcount4_test

import (
	"testing"

	popcountloop "gopl.io/ch2/ex2-3"

	popcount3 "gopl.io/ch2/ex2-4"
	popcount4 "gopl.io/ch2/ex2-5"
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

func BenchmarkPopCount3Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount3.PopCount(25534)
	}
}

func BenchmarkPopCount4Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount4.PopCount(25534)
	}
}
