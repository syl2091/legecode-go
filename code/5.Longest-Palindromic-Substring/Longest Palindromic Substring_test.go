package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := "lege"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}
