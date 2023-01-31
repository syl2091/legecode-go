package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var arr = []int{1, 2, 3}
	var arr1 = arr[:len(arr)-1]
	fmt.Printf("%v", arr1)
}

func TestLongestValidParentheses(t *testing.T) {
	int := longestValidParentheses("((()))()()()()()")
	fmt.Println(int)
}
