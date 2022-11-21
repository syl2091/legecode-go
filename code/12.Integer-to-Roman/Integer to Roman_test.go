package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	roman := intToRoman(1998)
	fmt.Println(roman)
}
