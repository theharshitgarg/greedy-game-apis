package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"project_greedy_game/db"
	"project_greedy_game/handlers"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Greedy Game Bidding and Auction Service")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/ad-slots", handlers.BookAdSlot)

	http.Handle("/", r)

	n := negroni.New()
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

    // Initializing the work
	db.GerInitialBidRequests()

	log.Print("Server Running on 127.0.0.1:8080")
	log.Fatal(srv.ListenAndServe())
}