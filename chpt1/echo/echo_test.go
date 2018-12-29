package main

import "testing"

// BenchmarkEcho1 benchmarks the echo1 function
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

// BenchmarkEcho2 benchmarks the echo2 function
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}
