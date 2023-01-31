package code

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "((())))"
	valid := isValid(s)
	fmt.Println(valid)
}
