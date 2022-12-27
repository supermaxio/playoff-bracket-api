package controller

import (
	"encoding/json"
	"net/http"

	"github.com/supermaxio/nflplayoffbracket/service"
)

func GetStandings(w http.ResponseWriter, r *http.Request) {
	standings, err := service.GetStandings()
	if err != nil {

	}

	json.NewEncoder(w).Encode(standings)
}
