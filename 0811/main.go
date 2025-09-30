package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, nex *ListNode
	pre, cur, nex = nil, head, head.Next
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = nex
		if nex != nil {
			nex = nex.Next
		}
	}
	return pre
}
