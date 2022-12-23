package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	rxgo "github.com/reactivex/rxgo/v2"
)

type Customer struct {
	ID, Age                   int
	Name, LastName, TaxNumber string
}

func main() {
	// Create an observable from a channel
	ch := make(chan rxgo.Item)
	go Producer(ch)
	observable := rxgo.FromChannel(ch)
	observable.ForEach(func(v interface{}) {
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})
	time.Sleep(1 * time.Minute)
	//
}

func Producer(ch chan rxgo.Item) {
	ch <- rxgo.Item{
		V: "Hello",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "everyone",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "!",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "I'm",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "a",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "stream",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "of",
	}
	time.Sleep(1 * time.Second)
	ch <- rxgo.Item{
		V: "strings",
	}
	///////////////////////////
	customerChan := make(chan rxgo.Item)
	go CustomerProducer(customerChan)

	observable := rxgo.FromChannel(customerChan)
	observable.Filter(func(v interface{}) bool {
		customer := v.(Customer)
		return customer.Age > 18
	}).Map(func(_ context.Context, item interface{}) (interface{}, error) {
		// Enrich operation
		customer := item.(Customer)
		fmt.Println("Assginning tax number to customer")
		customer.TaxNumber = RandomlyGenerateString()
		return customer, nil
	}).ForEach(func(v interface{}) {
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})
	time.Sleep(1 * time.Minute)
	///////////////////////////

}

func RandomlyGenerateString() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	return fmt.Sprintf("TN%d", n)
}

func CustomerProducer(
	ch chan rxgo.Item,
) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		age := rand.Intn(100)
		customer := Customer{
			ID:       i,
			Age:      age,
			Name:     fmt.Sprintf("Name %d", i),
			LastName: fmt.Sprintf("LastName %d", i),
		}
		ch <- rxgo.Item{
			V: customer,
		}
		time.Sleep(1 * time.Second)
	}
}
