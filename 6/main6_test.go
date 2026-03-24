package main

import "testing"

func TestRandNumGeneratorCount(t *testing.T) {
	count := 10
	max := 100

	var gotCount int
	for range randNumGenerator(count, max) {
		gotCount++
	}

	if gotCount != count {
		t.Errorf("randNumGenerator() count = %d, want %d", gotCount, count)
	}
}

func TestRandNumGeneratorRange(t *testing.T) {
	count := 20
	max := 100

	for num := range randNumGenerator(count, max) {
		if num < 0 || num >= max {
			t.Errorf("randNumGeneratorRange error")
		}
	}
}

func TestRandNumGeneratorZeroCount(t *testing.T) {
	var gotCount int
	for range randNumGenerator(0, 100) {
		gotCount++
	}

	if gotCount != 0 {
		t.Errorf("randNumGenerator() count = %d, want 0", gotCount)
	}
}
