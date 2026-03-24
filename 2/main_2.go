package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(30)
	}
	fmt.Println(originalSlice)
	fmt.Println(sliceExample(originalSlice))
	fmt.Println(addElements(originalSlice, 1))
	fmt.Println(copySlice(originalSlice))
	fmt.Println(removeElement(originalSlice, 9))
}

func sliceExample(slice []int) []int {
	newSlice := make([]int, 0)
	for _, v := range slice {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func addElements(slice []int, value int) []int {
	slice = append(slice, value)
	return slice
}

func copySlice(slice []int) []int {
	var newSlice []int
	newSlice = append(newSlice, slice...)

	return newSlice
}

var ErrInvalidIndex = errors.New("Неверный индекс")

func removeElement(slice []int, value int) ([]int, error) {
	if value < 0 || value >= len(slice) {
		return slice, ErrInvalidIndex
	}
	return slices.Delete(slice, value, value+1), nil
}
