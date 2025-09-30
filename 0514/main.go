package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	for p1 != nil {
		p.Next = p1
		p1 = p1.Next
	}
	for p2 != nil {
		p.Next = p2
		p2 = p2.Next
	}
	return dummy.Next

}

func main() {
	l1 := &ListNode{1,
		&ListNode{3,
			&ListNode{5, nil},
		},
	}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}
	l3 := mergeList(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
