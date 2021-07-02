package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// type FileSystemPlayerStore struct {
// 	database io.ReadWriteSeeker
// 	league   League
// }

const jsonContentType = "application/json"

// func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
// 	database.Seek(0, 0)
// 	league, _ := NewLeague(database)
// 	return &FileSystemPlayerStore{
// 		database: database,
// 		league:   league,
// 	}
// }

// func (f *FileSystemPlayerStore) GetLeague() League {
// 	return f.league
// }

// func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

// 	player := f.league.Find(name)

// 	if player != nil {
// 		return player.Wins
// 	}

// 	return 0
// }

// func (f *FileSystemPlayerStore) RecordWin(name string) {
// 	player := f.league.Find(name)

// 	if player != nil {
// 		player.Wins++
// 	} else {
// 		f.league = append(f.league, Player{name, 1})
// 	}

// 	f.database.Seek(0, 0)
// 	json.NewEncoder(f.database).Encode(f.league)
// }

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.Header().Set("content-type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
