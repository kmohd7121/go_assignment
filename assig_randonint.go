package main

import (
	"fmt"
	"math/rand"
)
func foo(c chan int, someValue int) {
	c <- rand.Intn(100 * someValue)

}

func main() {
	fooVal := make(chan int)
	go foo(fooVal,6)
	go foo(fooVal,2)
	v1 := <-fooVal
	v2 := <-fooVal
	fmt.Println(v1, v2)
}