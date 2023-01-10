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

	secure := router.PathPrefix("/v1/brackets").Subrouter()
	secure.Use(auth.JwtVerify)
	secure.HandleFunc("/", controller.BracketsController).Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
	secure.HandleFunc("/playoff_standings", controller.GetPlayoffStandings).Methods("GET", "OPTIONS")
	secure.HandleFunc("/refresh_standings", controller.GetStandings).Methods("GET", "OPTIONS")

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
