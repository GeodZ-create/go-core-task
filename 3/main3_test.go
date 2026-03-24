package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	s := StringIntMap{}
	s.Add("odin", 1)

	want := map[string]int{
		"odin": 1,
	}

	if !reflect.DeepEqual(s.data, want) {
		t.Errorf("Add() = %v, want %v", s.data, want)
	}

}

func TestRemove(t *testing.T) {
	s := StringIntMap{}
	s.data = map[string]int{
		"odin": 1,
		"dva":  2,
		"tri":  3,
	}
	s.Remove("dva")

	want := map[string]int{
		"odin": 1,
		"tri":  3,
	}
	if !reflect.DeepEqual(s.data, want) {
		t.Errorf("Remove() = %v, want %v", s.data, want)
	}

}

func TestCopy(t *testing.T) {
	s := StringIntMap{}
	s.data = map[string]int{
		"odin": 1,
		"dva":  2,
		"tri":  3,
	}
	copiedMap := s.Copy()

	if !reflect.DeepEqual(s.data, copiedMap) {
		t.Errorf("Copy() = %v, want %v", s.data, copiedMap)
	}
}

func TestExists(t *testing.T) {
	s := StringIntMap{}
	s.data = map[string]int{
		"odin": 1,
		"dva":  2,
		"tri":  3,
	}
	exist := s.Exists("dva")

	if !exist {
		t.Errorf("Exists() = %v, want %v", s.data, exist)
	}
}

func TestGet(t *testing.T) {
	s := StringIntMap{}
	s.data = map[string]int{
		"odin": 1,
		"dva":  2,
		"tri":  3,
	}
	get, ok := s.Get("dva")

	if get != 2 || ok != true {
		t.Errorf("Get() = %v, ok %v", get, ok)
	}
}

func TestGetInvalid(t *testing.T) {
	s := StringIntMap{}
	s.data = map[string]int{
		"odin": 1,
		"dva":  2,
		"tri":  3,
	}
	get, ok := s.Get("chetire")

	if get != 0 || ok != false {
		t.Errorf("Get() = %v, ok %v", get, ok)
	}
}
