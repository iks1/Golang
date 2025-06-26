package main 

import (
	"encoding/json"
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

const jsonContentType = "application/json"

//qs: we could have struct here instead of this interface right ? 
type PlayerStore interface{
	GetPlayerScore(name string) int
	RecordWin(name string)
    GetLeague() []Player
}

type PlayerServer struct{
	store PlayerStore 
	http.Handler
}

type Player struct {
	Name string 
	Wins int 
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store 

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leageHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

    p.Handler = router
	return p

}

func(p *PlayerServer) leageHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", jsonContentType	)
	json.NewEncoder(w).Encode(p.GetLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func(p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")

		switch r.Method {
		case http.MethodGet:
			p.showScore(w, r, player)
		case http.MethodPost:
			p.processWin(w, r, player)
		}
	
}

func (p *PlayerServer) GetLeagueTable() []Player {
	return p.store.GetLeague()
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