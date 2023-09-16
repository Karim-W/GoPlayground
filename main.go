package main

import (
	"fmt"
	"strconv"
)

func main() {
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func IntToString2(n int) string {
	return fmt.Sprintf("%d", n)
}

func IntToString3(n int) string {
	str := strconv.FormatInt(int64(n), 10)
	return str
}

/*
➜ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: GoPlayground
BenchmarkIntToString-10         46341439                24.69 ns/op           16 B/op          1 allocs/op
BenchmarkIntToString2-10        17687136                67.71 ns/op           24 B/op          2 allocs/op
BenchmarkIntToString3-10        48188415                24.65 ns/op           16 B/op          1 allocs/op
PASS
ok      GoPlayground    4.444s
*/
