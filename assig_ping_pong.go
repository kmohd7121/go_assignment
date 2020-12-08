package main

import (
	"fmt"
	"time"
	"os"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(200 * time.Millisecond)
		ponger <- 1
	}
}
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(100 * time.Millisecond)
		pinger <- 1
	}
}
func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)
	ping <- 1
	for{
		time.Sleep(time.Second)
	}
}
