package main

import (
	"fmt"
	"sync"
)

func joinChannels(chans ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chans))

		for _, ch := range chans {
			go func(ch <-chan int) {
				defer wg.Done()
				for val := range ch {
					mergedCh <- val
				}
			}(ch)

		}
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for _, num := range []int{4, 5, 6} {
			ch2 <- num
		}
		close(ch2)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			ch3 <- num
		}
		close(ch3)
	}()

	for num := range joinChannels(ch1, ch2, ch3) {
		fmt.Println(num)
	}
}
