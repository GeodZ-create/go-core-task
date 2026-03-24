package main

import (
	"fmt"
)

type semaphore chan struct{}

func NewSemaphore(n int) semaphore {
	return make(chan struct{}, n)
}

func (s semaphore) Done() {
	s <- struct{}{}
}
func (s semaphore) Wait(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	wg := NewSemaphore(5)
	go func() {
		defer wg.Done()
		a := 5
		fmt.Println(a)
	}()
	go func() {
		defer wg.Done()
		b := 2
		fmt.Println(b)
	}()
	go func() {
		defer wg.Done()
		c := 3
		fmt.Println(c)
	}()
	go func() {
		defer wg.Done()
		d := 1
		fmt.Println(d)
	}()
	go func() {
		defer wg.Done()
		e := 7
		fmt.Println(e)
	}()
	wg.Wait(5)

}
