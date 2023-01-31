package code

import (
	"fmt"
	"structures"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var arr = []int{1, 3, 4, 5, 6, 7, 8}
	var n = 10
	cycle := structures.Ints2List(arr)
	end := removeNthFromEnd(cycle, n)
	fmt.Printf("%v", end)
}
