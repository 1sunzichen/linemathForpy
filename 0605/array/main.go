package main

type MyArrayStack[T any] struct {
	list []T
}

func (s *MyArrayStack[T]) Push(e T) {
	s.list = append(s.list, e)
}

func (s *MyArrayStack[T]) Pop() T {
	if len(s.list) == 0 {
		var zero T
		return zero
	}
	e := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return e
}
