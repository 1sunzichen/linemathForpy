package main

import "container/list"

type MyLinkedStack struct {
	list *list.List
}

func NewMyLinkedStack() *MyLinkedStack {
	return &MyLinkedStack{
		list: list.New(),
	}
}
