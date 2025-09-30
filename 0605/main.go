package main

import (
	"container/list"
	"fmt"
)

type MyLinkedStack struct {
	list *list.List
}

func NewMyLinkedStack() *MyLinkedStack {
	return &MyLinkedStack{
		list: list.New(),
	}
}

func (s *MyLinkedStack) Push(val int) {
	s.list.PushBack(val)
}
func (s *MyLinkedStack) Pop() int {
	back := s.list.Back()
	if back != nil {
		s.list.Remove(back)
		return back.Value.(int)
	}
	return -1
}
func (s *MyLinkedStack) Peek() int {
	element := s.list.Back()
	if element != nil {
		return element.Value.(int)
	}
	return -1
}
func (s *MyLinkedStack) Size() int {
	return s.list.Len()
}

func main() {
	stack := NewMyLinkedStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
