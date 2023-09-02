package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("no")
}

func StringBuild() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func StringBuild2() string {
	var s string
	for i := 0; i < 1000; i++ {
		s = fmt.Sprintf("%sa", s)
	}
	return s
}

func StringBuild3() string {
	var s string
	for i := 0; i < 1000; i++ {
		s = s + "a"
	}
	return s
}

func StringBuild4() string {
	buf := make([]byte, 1000)
	bp := 0
	for i := 0; i < 1000; i++ {
		bp += copy(buf[bp:], "a")
	}
	return string(buf)
}

func StringBuild5() string {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	for i := 0; i < 1000; i++ {
		buf.WriteString("a")
	}
	return buf.String()
}

func StringBuild6() string {
	builder := strings.Builder{}
	for i := 0; i < 1000; i++ {
		builder.WriteString("a")
	}
	return builder.String()
}

func StringBuild7() string {
	buf := make([]byte, 0, 1000)
	for i := 0; i < 1000; i++ {
		buf = append(buf, "a"...)
	}
	return string(buf)
}
