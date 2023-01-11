package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supermaxio/nflplayoffbracket/customerrors"
	"github.com/supermaxio/nflplayoffbracket/service"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func UsersController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usernameToGet := vars["USERNAME"]
	switch r.Method {
	case "GET":
		if usernameToGet != "" {
			userToReturn, err := service.GetUser(usernameToGet)
			if err != nil {
				customerrors.HttpError(w, r, http.StatusBadRequest, "Unable to find user", err)
				return
			}

			json.NewEncoder(w).Encode(userToReturn)
		} else {

			usersToReturn, err := service.GetAllUsers()
			if err != nil {
				customerrors.HttpError(w, r, http.StatusBadRequest, "Unable to find users", err)
				return
			}

			json.NewEncoder(w).Encode(usersToReturn)
		}
	case "PUT":
		userUpdateModel := types.UserUpdate{}
		json.NewDecoder(r.Body).Decode(&userUpdateModel)
		userToReturn, err := service.UpdateUser(usernameToGet, userUpdateModel)
		if err != nil {
			customerrors.HttpError(w, r, http.StatusBadRequest, "Unable to find users", err)
			return
		}

		json.NewEncoder(w).Encode(userToReturn)
	}
}
