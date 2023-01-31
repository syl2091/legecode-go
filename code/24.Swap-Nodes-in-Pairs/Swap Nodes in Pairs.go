package code

import "structures"

type ListNode = structures.ListNode

func swapPairs(node *ListNode) *ListNode {
	dummy := &ListNode{Next: node}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}
