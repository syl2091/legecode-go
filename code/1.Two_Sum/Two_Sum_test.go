package code

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6}
	target := 10

	ints2 := twoSum(ints, target)

	fmt.Printf("%v", ints2)
}
