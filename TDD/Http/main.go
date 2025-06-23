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


func main(){
	// before PlayerServer is struct
	//handler := http.HandlerFunc(PlayerServer)
	
	// after PlayerServer is struct 
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}