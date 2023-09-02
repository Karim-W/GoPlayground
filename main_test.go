package main

import (
	"testing"
)

// run on cmd : go test -bench=. -benchmem
func BenchmarkAppendSpeed(b *testing.B) {
	appendSpeed(b.N)
}

func BenchmarkAppendSpeed2(b *testing.B) {
	appendSpeed2(b.N)
}

func BenchmarkAppendSpeed3(b *testing.B) {
	appendSpeed3(b.N)
}
