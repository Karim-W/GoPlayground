package main

import "testing"

func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToString(1234567890)
	}
}

func BenchmarkIntToString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToString2(1234567890)
	}
}

func BenchmarkIntToString3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToString3(1234567890)
	}
}
