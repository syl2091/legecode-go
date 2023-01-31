package main

import (
	"fmt"
	"testing"
)

func TestDictionary_Search(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	err, def := dictionary.Search("lege")
	if err != nil {
		panic(err)
	}
	fmt.Println(def)
}
