package main

import (
	"fmt"
	"runtime"
	"time"
)

var myarray []int

func memLeak() {
	temparr := make([]int, 999999)
	myarray = temparr[:5]
	fmt.Println(myarray)
}

func lessMemLeak() {
	temparr := make([]int, 999999)
	myarray = nil
	myarray = append(myarray, temparr[:5]...)
	fmt.Println(myarray)
}

func main() {
	fmt.Println("memleak")
	time.Sleep(time.Second)

	memLeak()

	runtime.GC()
	time.Sleep(time.Second)
	printMemUsage()
	time.Sleep(time.Second)
	fmt.Println("lessMemLeak")
	runtime.GC()
	time.Sleep(time.Second)

	lessMemLeak()
	runtime.GC()
	time.Sleep(time.Second)
	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	//   https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %0.2f MiB", bToMbyte(m.Alloc))
	fmt.Printf("\tTotalAlloc = %0.2f MiB", bToMbyte(m.TotalAlloc))
	fmt.Printf("\tSys = %0.2f MiB", bToMbyte(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMbyte(b uint64) float64 {
	return float64(b) / float64(1024) / float64(1024)
}

func appendSpeed(number int) {
	arryint := make([]int, 0)

	for i := 0; i < number; i++ {
		arryint = append(arryint, i)
	}
}

func appendSpeed2(number int) {
	arryint := make([]int, number)
	for i := 0; i < number; i++ {
		arryint[i] = i
	}
}

func appendSpeed3(number int) {
	arryint := make([]int, 0, number)

	for i := 0; i < number; i++ {
		arryint = append(arryint, i)
	}
}
