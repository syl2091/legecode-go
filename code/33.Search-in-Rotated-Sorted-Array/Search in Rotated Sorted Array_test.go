package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var arr = []int{4, 5, 6, 7, 0, 1, 2}

	i := search(arr, 7)
	fmt.Println(i)
}
