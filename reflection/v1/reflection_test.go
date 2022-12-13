package main

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	s := "lege"
	s1 := struct {
		name string
	}{s}

	var got []string

	walk(s1, func(s string) {
		fmt.Println(append(got, s))
	})
}
