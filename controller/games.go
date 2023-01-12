package controller

import (
	"encoding/json"
	"net/http"

	"github.com/supermaxio/nflplayoffbracket/customerrors"
	"github.com/supermaxio/nflplayoffbracket/service"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	standings, err := service.GetGames()
	if err != nil {
		customerrors.HttpError(w, r, http.StatusBadRequest, "error while getting games", err)
		return
	}

	json.NewEncoder(w).Encode(standings)
}

func RefreshScores(w http.ResponseWriter, r *http.Request) {
	standings, err := service.RefreshScores()
	if err != nil {
		customerrors.HttpError(w, r, http.StatusBadRequest, "error while refreshing scores", err)
		return

	}

	json.NewEncoder(w).Encode(standings)
}
