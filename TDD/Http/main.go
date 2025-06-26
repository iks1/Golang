package main 

import (
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayStore {
	return &InMemoryPlayStore{map[string]int{}}
}

type InMemoryPlayStore struct{
	store map[string]int
}

func (i *InMemoryPlayStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}


func main(){
	// before PlayerServer is struct
	//handler := http.HandlerFunc(PlayerServer)
	
	// after PlayerServer is struct 
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}