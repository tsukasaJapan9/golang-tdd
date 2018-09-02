package main

import (
	"net/http"
	"log"
	"fmt"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func main() {
	server := &PlayerServer{}

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen:%v", err)
	}
}