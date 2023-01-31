package main

import (
	"fmt"
	"testing"
)

func TestDictionary_Add(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	dictionary.Add("lege", "lege")
	dictionary.Add("lege", "lege1")
	fmt.Printf("%v", dictionary)
}
