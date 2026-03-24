package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestJoinChannels_Basic(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{4, 5, 6} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			c <- num
		}
		close(c)
	}()
	var channelSlice []int
	for num := range joinChannels(a, b, c) {
		channelSlice = append(channelSlice, num)
	}
	basicSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slices.Sort(channelSlice)

	if !reflect.DeepEqual(channelSlice, basicSlice) {
		t.Errorf("TestJoinChannelsBasic got = %v, want %v", channelSlice, basicSlice)
	}
}

func TestJoinChannels_EmptyChannel(t *testing.T) {
	a := make(chan int)
	b := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{} {
			b <- num
		}
		close(b)
	}()

	var channelSlice []int
	for num := range joinChannels(a, b) {
		channelSlice = append(channelSlice, num)
	}
	basicSlice := []int{1, 2, 3}
	slices.Sort(channelSlice)

	if !reflect.DeepEqual(channelSlice, basicSlice) {
		t.Errorf("TestJoinChannelsEmpty got = %v, want %v", channelSlice, basicSlice)
	}

}

func TestJoinChannels_AllEmpty(t *testing.T) {
	a := make(chan int)
	b := make(chan int)

	go func() {
		for _, num := range []int{} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{} {
			b <- num
		}
		close(b)
	}()

	var channelSlice []int
	for num := range joinChannels(a, b) {
		channelSlice = append(channelSlice, num)
	}
	var basicSlice []int
	slices.Sort(channelSlice)

	if !reflect.DeepEqual(channelSlice, basicSlice) {
		t.Errorf("TestJoinChannelsEmpty got = %v, want %v", channelSlice, basicSlice)
	}

}
