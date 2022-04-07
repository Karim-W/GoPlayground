package main

import (
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	go func() {
		// start your software here. Maybe your need to replace the for loop with other code
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run()
	}()
	c1 := CreateCallabeFunction(test)
	c2 := CreateCallabeFunction(intTwo, 2)
	HandleExits(c1, c2)
}

func test() {
	log.Println("i wanna die")
}
func intTwo(a int) int {
	log.Println(a * 2)
	return a * 2
}

type CallableFunction struct {
	fun  any
	args []reflect.Value
}

func CreateCallabeFunction(fun any, args ...any) CallableFunction {
	functionArguments := make([]reflect.Value, len(args))
	for i := range args {
		functionArguments[i] = reflect.ValueOf(args[i])
	}
	return CallableFunction{
		fun:  fun,
		args: functionArguments,
	}
}

func HandleExits(args ...CallableFunction) {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan
	log.Printf("Caught SIGTERM %v", sig)
	log.Println("Exiting...")
	for i := range args {
		fc := args[i]
		v := reflect.ValueOf(fc.fun)
		if v.Kind() != reflect.Func {
			panic("argument must be a function")
		} else {
			v.Call(fc.args)
		}

	}

	//Something like this?
	//workyy
	//Need to clean up though but yay
	//yup yup
	//i go back to pachirisu/zogor now. bye
	// lettuce see
	//
	// for i := 0; i < 10; i++ {
	// 	log.Println("i wanna die")
	// }

}
