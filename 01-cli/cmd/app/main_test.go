package main

import "testing"

func BenchmarkWriteCliArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeCliArgs()
	}
}
