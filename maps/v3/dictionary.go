package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}

func (d Dictionary) Add(word, def string) {
	d[word] = def
}
