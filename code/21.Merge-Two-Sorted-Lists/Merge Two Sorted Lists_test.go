package code

import (
	"fmt"
	"structures"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var arr1 = []int{1, 2, 3}
	var arr2 = []int{1, 3, 4}
	list1 := structures.Ints2List(arr1)
	list2 := structures.Ints2List(arr2)
	listNode := mergeTwoLists(list1, list2)
	fmt.Printf("%v", structures.List2Ints(listNode))
}
