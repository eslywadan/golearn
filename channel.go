package main

import (
	"fmt"
	"math"
	"time"
)

func work(workCh chan int, resultCh chan float64) {
	// fetch from work channel
	x := <-workCh

	time.Sleep(1 * time.Second)

	resultCh <- math.Sqrt(float64(x))
}

func main() {
	// unbuffered channel
	//
	// note that read/write with same unbuffered channel may meed to be run 
	// in different goroutines to prevent deadlock
	//
	// use unbuffered channel carefully
	workCh := make(chan int)
	
	// buffered channel
	//
	// buffered channel will block read/write if full
	resultCh := make(chan float64, 10)

	for i := 0; i < 10; i++ {
		go work(workCh, resultCh)
	}

	for i := 0; i < 10; i++ {
		workCh <- i
	}

	for i :=0; i< 10; i++ {
		fmt.Println(<-resultCh)
	} 
}