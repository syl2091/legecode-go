package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("can not found the word")

func (d Dictionary) Search(word string) (error, string) {
	def, ok := d[word]
	if !ok {
		return ErrNotFound, ""
	}
	return nil, def
}
