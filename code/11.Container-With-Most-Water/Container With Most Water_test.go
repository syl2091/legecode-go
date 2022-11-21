package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ints := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(ints)
	fmt.Println(area)
}
