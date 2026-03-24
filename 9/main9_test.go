package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestPipelineBasic(t *testing.T) {
	var a uint8 = 2
	var b uint8 = 3
	c := FirstChan(a, b)

	var got []float64

	want := []float64{8.0, 27.0}

	for out := range Cubing(c) {
		got = append(got, out)
	}
	slices.Sort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestPipelineBasic want = %v, got %v", want, got)

	}
}

func TestPipelineEmpty(t *testing.T) {
	c := FirstChan()

	got := []float64{}

	want := []float64{}

	for out := range Cubing(c) {
		got = append(got, out)
	}
	slices.Sort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestPipelineEmpty want = %v, got %v", want, got)
	}
}
