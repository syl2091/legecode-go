package main

import (
	"log"
	"net/http"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestInMemoryPlayerStore_GetPlayerScore(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"lege":  20,
		"lege1": 30,
	}}
	server := &PlayerServer{&store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
