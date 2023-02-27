package common

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	word := SaltPassWord("syl2091", "lege")
	fmt.Println(word)
}
