package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{1, 3}

	ints := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ints)
}
