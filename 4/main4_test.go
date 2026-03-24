package main

import (
	"reflect"
	"testing"
)

func TestNewSlice_BasicCase(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := NewSlice(slice1, slice2)

	want := []string{"apple", "cherry", "43", "lead", "gno1"}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("NewSlice() = %v, want %v", result, want)
	}
}

func TestNewSlice_NoMatches(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"x", "y"}
	result := NewSlice(slice1, slice2)

	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("NewSlice() = %v, want %v", result, want)
	}
}

func TestNewSlice_AllElementsExcluded(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{"a", "b", "c"}
	result := NewSlice(slice1, slice2)

	want := []string{}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("NewSlice() = %v, want %v", result, want)
	}
}
