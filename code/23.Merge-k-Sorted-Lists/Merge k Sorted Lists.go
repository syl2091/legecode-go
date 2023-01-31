package code

import "structures"

type ListNode = structures.ListNode

func mergeKLists(nodes []*ListNode) *ListNode {
	len := len(nodes)
	if len < 1 {
		return nil
	}
	if len == 1 {
		return nodes[0]
	}
	num := len / 2
	left := mergeKLists(nodes[:num])
	right := mergeKLists(nodes[num:])
	return mergeTwoLists1(left, right)
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}
