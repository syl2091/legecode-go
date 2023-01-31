package code

import (
	"fmt"
	"structures"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	list := structures.Ints2List([]int{1, 2, 3, 4})
	list1 := swapPairs(list)
	fmt.Printf("%v", structures.List2Ints(list1))
}
