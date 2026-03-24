package main_1

import (
	"reflect"
	"testing"
)

func TestGetType(t *testing.T) {
	got := getType(1)

	want := "int"

	if got != want {
		t.Errorf("getType() = %q, want %q", got, want)
	}
}

func TestValueToString(t *testing.T) {
	got := valueToString(1, 1.0, "asd", true, complex64(0+0i))
	want := "11asdtrue(0+0i)"

	if got != want {
		t.Errorf("valueToString() = %q, want %q", got, want)
	}
}

func TestStringToRune(t *testing.T) {
	got := stringToRune("Go")

	want := []rune{'G', 'o'}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("stringToRune() = %v, want = %v", got, want)
	}
}

func TestPutSalt(t *testing.T) {
	got := putSalt([]rune("go"))
	want := []rune("ggo-2024o")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("putSalt() = %v, want %v", got, want)
	}
}

func TestGetHash(t *testing.T) {
	got := getHash([]rune("test"))
	want := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"

	if got != want {
		t.Errorf("getHash() = %q, want %q", got, want)
	}

}
