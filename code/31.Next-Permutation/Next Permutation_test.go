package code

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var arr = []int{3, 2, 1, 4, 5, 6}
	nextPermutation(arr)
	fmt.Printf("%v", arr)
}
