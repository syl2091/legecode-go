package main

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	x := struct {
		name     string
		birthday string
		age      string
	}{"lege",
		"19980619",
		"25"}
	var got []string

	walk(x, func(input string) {
		fmt.Println(append(got, input))
	})

}
