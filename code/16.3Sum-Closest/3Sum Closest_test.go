package code

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	closest := threeSumClosest([]int{0, 0, 0}, 1)
	fmt.Printf("%v", closest)
}

func TestName2(t *testing.T) {
	ints := make([]int, 100)
	for i := range ints {
		ints = append(ints, i)
	}
	closest := threeSumClosest(ints, 300)
	closest = closest
	fmt.Printf("%v", closest)
}
