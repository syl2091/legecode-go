package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerStore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
