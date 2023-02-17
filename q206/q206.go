package q206

import "github.com/konomichael/leetcode/datastructure"

type ListNode = datastructure.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}

	return pre
}
