package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/supermaxio/nflplayoffbracket/auth"
	"github.com/supermaxio/nflplayoffbracket/customerrors"
	"github.com/supermaxio/nflplayoffbracket/service"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func BracketsController(w http.ResponseWriter, r *http.Request) {
	bracket := types.Bracket{}
	username := auth.GetClaims().Username
	switch r.Method {
	case "GET":
		bracketToReturn, err := service.GetBracket(username)
		if err != nil {
			customerrors.HttpError(w, r, http.StatusBadRequest, fmt.Sprintf("Unable to find bracket for %s", username), err)
			return
		}

		json.NewEncoder(w).Encode(bracketToReturn)
	case "POST":
		json.NewDecoder(r.Body).Decode(&bracket)

		bracket.Username = username
		log.Println(bracket)
		bracketToReturn, err := service.CreateBracket(bracket)
		if err != nil {
			customerrors.HttpError(w, r, http.StatusBadRequest, fmt.Sprintf("Unable to create bracket for %s", username), err)
			return
		}

		json.NewEncoder(w).Encode(bracketToReturn)
	case "PUT":
		json.NewDecoder(r.Body).Decode(&bracket)

		bracket.Username = username
		bracketToReturn, err := service.UpdateBracket(bracket)
		if err != nil {
			customerrors.HttpError(w, r, http.StatusBadRequest, err.Error(), err)
			return
		}

		json.NewEncoder(w).Encode(bracketToReturn)
	case "DELETE":
		err := service.DeleteBracket(username)
		if err != nil {
			customerrors.HttpError(w, r, http.StatusBadRequest, fmt.Sprintf("Unable to delete bracket for %s", username), err)
			return
		}

		json.NewEncoder(w).Encode(fmt.Sprintf("Deleted bracket of user: %s", username))
	}
}
