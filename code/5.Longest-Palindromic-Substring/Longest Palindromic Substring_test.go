package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := "lege"
	palindrome := longestPalindrome1(s)
	fmt.Println(palindrome)
}
