package code

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	strings := generateParenthesis(4)
	fmt.Printf("%v", strings)
}
