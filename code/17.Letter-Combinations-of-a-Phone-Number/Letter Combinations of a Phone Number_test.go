package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	strings := letterCombinations("456")
	fmt.Printf("%v", strings)
}
