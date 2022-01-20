package main

import (
	"fmt"
	"sync"

	// "sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Rout() {
	tStart := time.Now().Unix()
	fmt.Println("Start:", tStart)
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "a", "b", "c", "d", "k", "l", "m", "n"}
	sliceLength := len(slice)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running for loop…")
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			val := slice[i]
			bcrypt.GenerateFromPassword([]byte(val+val+val+val+val+val+val+val+val+val+val+val+val+val), bcrypt.DefaultCost)
			fmt.Printf("i: %v, val: %v\n", i, val)
		}(i)
	}
	wg.Wait()
	// fmt.Println("Running for loop…")
	// for i := 0; i < sliceLength; i++ {
	// 	val := slice[i]
	// 	bcrypt.GenerateFromPassword([]byte(val+val+val+val+val+val+val+val+val+val+val+val+val+val), bcrypt.DefaultCost)
	// 	fmt.Printf("i: %v, val: %v\n", i, val)
	// }
	// fmt.Println("Finished for loop")
	fmt.Println("Finished for loop")
	tEnd := time.Now().Unix()
	fmt.Println("End:", tEnd)
	fmt.Printf("Time taken: %v\n", tEnd-tStart)

}
