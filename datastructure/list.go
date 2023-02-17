package datastructure

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func IsListEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a, b = a.Next, b.Next
	}
	if a == nil && b == nil {
		return true
	}

	return false
}

func MakeListNode(vals ...int) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead

	for _, val := range vals {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return dummyHead.Next
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}

	vals := []string{}
	for n != nil {
		vals = append(vals, strconv.Itoa(n.Val))
		n = n.Next
	}
	return strings.Join(vals, "->")
}
