package code

import (
	"fmt"
	"structures"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var arr1 = []int{1, 4, 5}
	var arr2 = []int{1, 3, 4}
	var arr3 = []int{2, 6}

	list1 := structures.Ints2List(arr1)
	list2 := structures.Ints2List(arr2)
	list3 := structures.Ints2List(arr3)

	nodes := make([]*ListNode, 0)
	listNodes := append(append(append(nodes, list1), list2), list3)
	lists := mergeKLists(listNodes)
	ints := structures.List2Ints(lists)
	fmt.Printf("%v", ints)
}
