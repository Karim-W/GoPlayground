package main

import (
	"GoPlayground/ginproto"
	"encoding/json"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
)

var sampleByte = []byte{10, 4, 116, 101, 115, 116, 16, 1}

func SampleHandler(res any) {
	var req ginproto.Request

	err := proto.Unmarshal(sampleByte, &req)
	if err != nil {
		panic(err)
	}

	resp := map[string]interface{}{}
	resp["message"] = sampleService(context.Background(), &req)
	*res.(*map[string]interface{}) = resp
}

func sampleService(ctx context.Context, req *ginproto.Request) string {
	return sampleRepository(ctx, req)
}

func sampleRepository(ctx context.Context, req *ginproto.Request) string {
	return sampleCore(ctx, req)
}

func sampleCore(ctx context.Context, req *ginproto.Request) string {
	return req.GetName()
}

type jsonReq struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

var sampleByteJSON = []byte(`{"name":"test","age":1}`)

func SampleHandlerJSON(res any) {
	var req jsonReq

	err := json.Unmarshal(sampleByteJSON, &req)
	if err != nil {
		panic(err)
	}

	resp := map[string]interface{}{}
	resp["message"] = sampleServiceJSON(context.Background(), req)
	*res.(*map[string]interface{}) = resp
}

func sampleServiceJSON(ctx context.Context, req jsonReq) string {
	return sampleRepositoryJSON(ctx, req)
}

func sampleRepositoryJSON(ctx context.Context, req jsonReq) string {
	return sampleCoreJSON(ctx, req)
}

func sampleCoreJSON(ctx context.Context, req jsonReq) string {
	return req.Name
}

func SampleHandlerHybrid(res any) {
	var req ginproto.Request

	err := proto.Unmarshal(sampleByte, &req)
	if err != nil {
		panic(err)
	}

	jsonReq := jsonReq{
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	resp := map[string]interface{}{}
	resp["message"] = sampleServiceJSON(context.Background(), jsonReq)
	*res.(*map[string]interface{}) = resp
}
