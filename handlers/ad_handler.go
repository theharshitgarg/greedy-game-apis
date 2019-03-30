package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
    "strconv"
	"github.com/gorilla/mux"

	"project_greedy_game/db"
	. "project_greedy_game/models"
)

func GetBidsOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	keys := r.URL.Query()
	fmt.Println("KEN", keys["bidId"])
	fmt.Fprintf(w, "Bid Category: %v\n", vars["bidId"], vars["book"])
}

// func BookAdSlot implements 1st part of the challenge.
func BookAdSlot(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

    if (len(keys["adId"]) < 1) {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
 	
    bId, _ := strconv.ParseInt(keys["adId"][0], 10, 64)
	isBook := false
	if keys["isBook"][0] == "true" {
		isBook = true

		bidR := BidRequest{}
		bidR.Create(bId, isBook)
		db.BidRequests = append(db.BidRequests, bidR)

		adtv := Avdertisement{}
		adtv.CreateAdvt(bidR.SID)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(adtv)

	} else {
		w.WriteHeader(http.StatusNoContent)
	}
	return
}
