package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//Counter displays the numbers from 0 do 100
func Counter() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Counter()
	}
	wg.Wait()
}
