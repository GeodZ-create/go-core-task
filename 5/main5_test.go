package main

import (
	"reflect"
	"testing"
)

func TestBasicCase(t *testing.T) {
	a := []int{65, 3, 57, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, result := MatchesSlice(a, b)

	want := []int{3, 64}

	if !reflect.DeepEqual(result, want) || !ok {
		t.Errorf("BasicCaseTest is bad result:%v want: %v", result, want)
	}
}

func TestNoMatches(t *testing.T) {
	a := []int{65, 3, 57}
	b := []int{63, 2, 53}

	ok, result := MatchesSlice(a, b)

	want := []int{}

	if !reflect.DeepEqual(result, want) || ok {
		t.Errorf("NoMatchesTest is bad result: %v, want: %v", result, want)
	}
}

func TestDublicateInFirstSlice(t *testing.T) {
	a := []int{65, 3, 3, 3, 2}
	b := []int{3, 2}

	ok, result := MatchesSlice(a, b)

	want := []int{3, 2}

	if !reflect.DeepEqual(result, want) || !ok {
		t.Errorf("DublicateInFirstSliceTest is bad result: %v, want: %v", result, want)
	}
}
