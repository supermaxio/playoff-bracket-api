package controller

import (
	"encoding/json"
	"net/http"

	"github.com/supermaxio/nflplayoffbracket/customerrors"
	"github.com/supermaxio/nflplayoffbracket/service"
)

func GetPlayoffStandings(w http.ResponseWriter, r *http.Request) {
	standings, err := service.GetPlayoffStandings()
	if err != nil {
		customerrors.HttpError(w, r, http.StatusBadRequest, "error while getting playoff standings", err)
		return
	}

	json.NewEncoder(w).Encode(standings)
}

func GetStandings(w http.ResponseWriter, r *http.Request) {
	standings, err := service.GetStandings()
	if err != nil {
		customerrors.HttpError(w, r, http.StatusBadRequest, "error while getting standings", err)
		return
	}

	json.NewEncoder(w).Encode(standings)
}
