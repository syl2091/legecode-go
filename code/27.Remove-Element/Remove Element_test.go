package code

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var arr = []int{3, 2, 2, 3}
	element := removeElement(arr, 3)
	fmt.Printf("%v", element)
}
