package main

import (
	"fmt"
	"time"
)

func Producer(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func Consumer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int, 1)

	go Producer(ch1)
	Consumer(ch1)
}
