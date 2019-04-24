package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Print prints Hello 10 times, one Hello at a second
func Print() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go Print()
	wg.Wait()
}
