package main

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	x := struct {
		name string
		x    interface{}
	}{"lege", struct {
		birthday string
		age      string
	}{"19980619",
		"25"}}
	var got []string
	walk(x.x, func(input string) {
		got = append(got, input)
	})
	fmt.Println(got)
}
