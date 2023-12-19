package main

import (
	"GoPlayground/ginproto"
	"fmt"
	"testing"

	"github.com/karim-w/stdlib/httpclient"
	"google.golang.org/protobuf/proto"
)

func TestClient(t *testing.T) {
	msg := ginproto.Request{Name: "test", Age: 1}

	body, err := proto.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(body)

	resp := httpclient.Req("http://localhost:8080").
		AddHeader("Content-Type", "application/x-protobuf").
		AddBodyRaw(body).
		Post()

	if !resp.IsSuccess() {
		t.Fatal(resp.CatchError())
	}

	t.Log(string(resp.GetBody()))
}

func BenchmarkProtoCallFlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SampleHandler(&map[string]interface{}{})
	}
}

func BenchmarkJSONCallFlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SampleHandlerJSON(&map[string]interface{}{})
	}
}

func BenchmarkHybridCallFlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SampleHandlerHybrid(&map[string]interface{}{})
	}
}

/*
goos: darwin
goarch: arm64
pkg: GoPlayground
BenchmarkProtoCallFlow-10        4478205               267.1 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4480280               266.6 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4426376               266.9 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4496472               266.7 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4489028               266.8 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4483308               266.7 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4508173               267.6 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4521552               266.0 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4524321               266.7 ns/op           420 B/op          5 allocs/op
BenchmarkProtoCallFlow-10        4498876               265.9 ns/op           420 B/op          5 allocs/op
BenchmarkJSONCallFlow-10         1830912               657.0 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1825809               656.9 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1825299               655.4 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1824297               659.5 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1826492               656.9 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1822632               657.8 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1823635               656.7 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1825978               656.5 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1825716               657.2 ns/op           600 B/op          9 allocs/op
BenchmarkJSONCallFlow-10         1825474               656.7 ns/op           600 B/op          9 allocs/op
BenchmarkHybridCallFlow-10       4500391               266.3 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4487716               266.0 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4486434               267.1 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4497476               266.9 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4489774               270.2 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4519418               266.1 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4498584               266.5 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4501641               268.1 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4508107               266.2 ns/op           420 B/op          5 allocs/op
BenchmarkHybridCallFlow-10       4463444               266.5 ns/op           420 B/op          5 allocs/op
PASS
ok      GoPlayground    49.178s

*/
