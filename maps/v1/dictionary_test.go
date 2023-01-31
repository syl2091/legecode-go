package main

import (
	"fmt"
	"testing"
)

func TestDictionary_Search(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	fmt.Println(got)
}
