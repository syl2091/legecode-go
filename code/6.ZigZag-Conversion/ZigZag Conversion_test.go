package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := convert("PAYPALISHIRING", 4)
	fmt.Println(s)
}
