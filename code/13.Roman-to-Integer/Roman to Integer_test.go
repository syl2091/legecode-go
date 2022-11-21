package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	toInt := romanToInt("MCMXCVIII")
	fmt.Println(toInt)
}
