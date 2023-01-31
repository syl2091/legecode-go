package code

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	ints := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Printf("%v", ints)
}
