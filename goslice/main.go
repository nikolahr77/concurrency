package main

import (
	"fmt"
	"time"
)

type Slice struct {
	slice []int
}

//Filler fills a slice with the numbers from 0 to 10
func (s *Slice) Filler() {
	for i := 0; i <= 100; i++ {
		s.slice = append(s.slice, i)
	}
}

//Filler1 fills a slice with the numbers from 0 to 10
func (s *Slice) Filler1() {
	for i := 110; i <= 200; i++ {
		s.slice = append(s.slice, i)
	}
}

//Filler2 fills a slice with the numbers from 0 to 10
func (s *Slice) Filler2() {
	for i := 210; i <= 300; i++ {
		s.slice = append(s.slice, i)
	}
}

//Filler3 fills a slice with the numbers from 0 to 10
func (s *Slice) Filler3() {
	for i := 310; i <= 400; i++ {
		s.slice = append(s.slice, i)
	}
}

func main() {
	s := Slice{}
	s.Filler()
	go s.Filler1()
	go s.Filler2()
	go s.Filler3()
	time.Sleep(1 * time.Second)
	fmt.Print(s)
}
