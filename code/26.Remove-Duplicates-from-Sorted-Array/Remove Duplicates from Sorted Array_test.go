package code

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := removeDuplicates(arr)
	fmt.Println(duplicates)
}
