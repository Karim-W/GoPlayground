package main

import "testing"

func BenchmarkStringBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild()
	}
}

func BenchmarkStringBuild2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild2()
	}
}

func BenchmarkStringBuild3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild3()
	}
}

func BenchmarkStringBuild4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild4()
	}
}

func BenchmarkStringBuild5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild5()
	}
}

func BenchmarkStringBuild6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild6()
	}
}

func BenchmarkStringBuild7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuild7()
	}
}

/*
Output:
➜ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: GoPlayground
BenchmarkStringBuild-10            15277             70671 ns/op          530278 B/op        999 allocs/op
BenchmarkStringBuild2-10            9247            123978 ns/op          546611 B/op       1998 allocs/op
BenchmarkStringBuild3-10           17454             67979 ns/op          530278 B/op        999 allocs/op
BenchmarkStringBuild4-10          346232              3433 ns/op            1024 B/op          1 allocs/op
BenchmarkStringBuild5-10          257548              4632 ns/op            1024 B/op          1 allocs/op
BenchmarkStringBuild6-10          412120              2852 ns/op            3320 B/op          9 allocs/op
BenchmarkStringBuild7-10         2266114               529.9 ns/op          1024 B/op          1 allocs/op
PASS
ok      GoPlayground    10.725s
*/
