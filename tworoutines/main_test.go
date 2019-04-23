package main

import "testing"

func TestChan(t *testing.T) {
	ch1 := make(chan int, 1)
	go Producer(ch1)
	Consumer(ch1)

}
