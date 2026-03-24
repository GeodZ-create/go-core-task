package main

import "maps"

func main() {

}

type StringIntMap struct {
	data map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	if s.data == nil {
		s.data = make(map[string]int)
	}
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	maps.Copy(newMap, s.data)
	return newMap

}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.data[key]
	return ok

}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.data[key]
	return v, ok
}
