package main

import "fmt"

func main() {
	timeNow := getTimeInISOString()
	fmt.Println(timeNow)
	fmt.Println(get30SecondsFromTime(timeNow))
}
