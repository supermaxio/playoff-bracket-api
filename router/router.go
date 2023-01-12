package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supermaxio/nflplayoffbracket/auth"
	"github.com/supermaxio/nflplayoffbracket/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.Use(auth.CorsHandler)
	// Add your routes as needed
	router.HandleFunc("/health", healthCheck)

	router.HandleFunc("/v1/register", auth.RegisterHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/login", auth.LoginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/refresh", auth.RefreshHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/logout", auth.Logout).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/scores/refresh", controller.RefreshScores).Methods("GET", "OPTIONS")

	standingsRouter := router.PathPrefix("/v1/standings").Subrouter()
	standingsRouter.Use(auth.JwtVerify)
	standingsRouter.HandleFunc("/playoffs", controller.GetPlayoffStandings).Methods("GET", "OPTIONS")
	standingsRouter.HandleFunc("/refresh", controller.GetStandings).Methods("GET", "OPTIONS")

	gamesRouter := router.PathPrefix("/v1/games").Subrouter()
	gamesRouter.Use(auth.JwtVerify)
	gamesRouter.HandleFunc("/", controller.GetGames).Methods("GET", "OPTIONS")

	bracketsRouter := router.PathPrefix("/v1/brackets").Subrouter()
	bracketsRouter.Use(auth.JwtVerify)
	bracketsRouter.HandleFunc("/", controller.BracketsController).Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
	bracketsRouter.HandleFunc("/{USERNAME}", controller.BracketsController).Methods("GET", "OPTIONS")

	usersRouter := router.PathPrefix("/v1/users").Subrouter()
	usersRouter.Use(auth.JwtVerify)
	usersRouter.HandleFunc("/", controller.UsersController).Methods("GET", "OPTIONS")
	usersRouter.HandleFunc("/{USERNAME}", controller.UsersController).Methods("GET", "PUT", "OPTIONS")
	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
