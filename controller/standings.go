package controller

import (
	"encoding/json"
	"net/http"

	"github.com/supermaxio/nflplayoffbracket/service"
)

func GetStandings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	standings, err := service.GetStandings()
	if err != nil {

	}

	json.NewEncoder(w).Encode(standings)
}
