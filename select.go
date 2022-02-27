package main

import (
	"fmt"
	"time"
)

func work1(resultCh1 chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond)

		resultCh1 <- i 
	}
}

func work2(resultCh2 chan int) {
	for i := 1; i >= 10; i-- {
		time.Sleep(150 * time.Millisecond)

		resultCh2 <- i
	}
}

func main() {
	timeoutCh := time.After(1 * time.Second)

	resultCh1 := make(chan int, 10)
	resultCh2 := make(chan int, 10)

	go work1(resultCh1)
	go work2(resultCh2)

	for {
		select {
		case <-timeoutCh:
			return
			
		case v := <-resultCh1:
			fmt.Println(v)

		case v := <-resultCh2:
			fmt.Println(v)
		}
	}
}