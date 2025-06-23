package main 

import (
	"fmt"
	"net/http"
	"strings"
)

// func GetPlayerScore(player string) string {
// 	if player == "pepper" {
// 		return "20"
// 	}

// 	if player == "floyd" {
// 		return "10"
// 	}
// 	return ""
// }

//qs: we could have struct here instead of this interface right ? 
type PlayerStore interface{
	GetPlayerScore(name string) int
	RecordWin(name string) 
}

type PlayerServer struct{
	store PlayerStore 
}

func (p* PlayerServer) ServeHTTP(w http.ResponseWriter, r* http.Request){
    
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r, player)
	case http.MethodPost:
		p.processWin(w, r, player)
	}
	
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request, player string){
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request, player string){
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}