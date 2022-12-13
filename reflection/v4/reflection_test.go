package main

import (
	"fmt"
	"testing"
)

type infomation struct {
	birthday string
	age      string
}

func TestWalk(t *testing.T) {
	x := struct {
		name  string
		input interface{}
	}{"lege", infomation{
		"19980619",
		"25",
	}}
	var got []string
	walk(x.input, func(input string) {
		got = append(got, input)
	})
	fmt.Printf("%v", got)
}

func TestWalk2(t *testing.T) {
	x := struct {
		name  string
		input interface{}
	}{"lege", struct {
		birthday string
		age      string
	}{
		"19980619",
		"25",
	}}
	walk2(x, func() {
	})
}
