package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(prefix)
}
