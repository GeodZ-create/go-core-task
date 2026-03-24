package main

import (
	"fmt"
	"slices"
)

func FirstChan(nums ...uint8) <-chan uint8 {
	out := make(chan uint8)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func Cubing(in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		for n := range in {
			out <- float64(n * n * n)
		}
		close(out)
	}()
	return out
}

func main() {
	c := FirstChan(2, 3, 4, 5, 6, 7)
	var result []float64
	for out := range Cubing(c) {
		result = append(result, out)
	}
	slices.Sort(result)
	fmt.Println(result)

}
