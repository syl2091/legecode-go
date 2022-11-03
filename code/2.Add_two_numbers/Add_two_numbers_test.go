package code

import (
	"fmt"
	"structures"
	"testing"
)

type param struct {
	one []int
	two []int
}

func TestAddTwoNumbers(t *testing.T) {
	param := &param{one: []int{1, 2, 3, 4}, two: []int{1, 2, 3, 4}}
	node1 := structures.Ints2List(param.one)
	//node2 := structures.Ints2List(param.two)
	ints := structures.List2Ints(AddTwoNumbers(node1, nil))
	fmt.Printf("%v", ints)

}
