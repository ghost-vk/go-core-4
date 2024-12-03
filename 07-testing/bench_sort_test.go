package main

import (
	"math/rand"
	"sort"
	"testing"
)

// 0.06645 ns/op
func Benchmark_Ints(b *testing.B) {
	l := 1_000_000
	ints := make([]int, l)
	for i := 0; i < l-1; i++ {
		ints[i] = rand.Int()
	}
	b.ResetTimer()
	sort.Ints(ints)
}

// 0.08701 ns/op
func Benchmark_Floats(b *testing.B) {
	l := 1_000_000
	floats := make([]float64, l)
	for i := 0; i < l-1; i++ {
		floats[i] = float64(rand.Float64())
	}
	b.ResetTimer()
	sort.Float64s(floats)
}
