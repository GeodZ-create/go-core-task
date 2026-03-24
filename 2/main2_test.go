package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	got := sliceExample([]int{18, 24, 18, 23, 1, 16, 7, 23, 6, 11})

	want := []int{18, 24, 18, 16, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample() = %v, want %v", got, want)
	}
}

func TestAddElements(t *testing.T) {
	got := addElements([]int{18, 24, 18, 23, 1, 16, 7, 23, 6, 11}, 1)

	want := []int{18, 24, 18, 23, 1, 16, 7, 23, 6, 11, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements() = %v, want %v", got, want)
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{18, 24, 18, 23, 1, 16, 7, 23, 6, 11}
	got := copySlice(original)
	original = append(original, 1)
	want := []int{18, 24, 18, 23, 1, 16, 7, 23, 6, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("copySlice() = %v, want %v", got, want)
	}
	if reflect.DeepEqual(got, original) {
		t.Errorf("Изменения оригинала влияют на копию")
	}
}

func TestRemoveElement(t *testing.T) {
	got, err := removeElement([]int{18, 24, 18, 23, 1, 16, 7, 23, 6, 11}, 0)
	if err != nil {
		t.Fatal("Ошибка удаления элемента")
	}

	want := []int{24, 18, 23, 1, 16, 7, 23, 6, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement() = %v, want %v", got, want)
	}
}

func TestRemoveElementInvalidIndex(t *testing.T) {
	original := []int{18, 24, 18, 23}
	got, err := removeElement(original, 4)

	if !errors.Is(err, ErrInvalidIndex) {
		t.Errorf("removeElement() = %v, want %v", got, original)
	}

	if !reflect.DeepEqual(got, original) {
		t.Errorf("removeElement() = %v, want %v", got, original)
	}
}
