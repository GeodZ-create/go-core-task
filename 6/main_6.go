package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumGenerator(count int, max int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- r.Intn(max)
		}
		close(out)
	}()
	return out
}

func main() {
	for num := range randNumGenerator(10, 100) {
		fmt.Println(num)
	}
}
