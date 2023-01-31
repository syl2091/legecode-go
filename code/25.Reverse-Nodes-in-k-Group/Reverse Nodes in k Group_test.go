package code

import (
	"fmt"
	"structures"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	node := structures.Ints2List(arr)
	kGroup := reverseKGroup(node, 3)
	fmt.Printf("%v", structures.List2Ints(kGroup))
}

func TestName(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	ints := structures.Ints2List(arr)
	node := reversek(ints, 2)
	arr1 := structures.List2Ints(node)
	fmt.Printf("%v", arr1)
}

func reversek(node *ListNode, k int) *ListNode {
	head := node
	for i := 0; i < k; i++ {
		if node == nil {
			return node
		}
		head = head.Next
	}
	return node
}
