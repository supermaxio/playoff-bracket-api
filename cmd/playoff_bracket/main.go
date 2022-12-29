package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/supermaxio/nflplayoffbracket/config"
	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/router"
)

func main() {
	// Setup configurations
	config.Setup()

	// Ping db
	database.MongoConnect()

	// Set up api router
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := router.Router()

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%s", config.GetDomain(), config.GetPort()),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Starting API at: " + server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	sigterm := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(sigterm, os.Interrupt)
	signal.Notify(sigterm, syscall.SIGTERM)

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Wait for the SIGTERM signal or the context to expire
	select {
	case <-sigterm:
		log.Println("Received SIGTERM, shutting down...")
	case <-ctx.Done():
		log.Println("Timeout exceeded, shutting down...")
	}

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// go srv.Shutdown(ctx)
	// Shut down the HTTP server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	database.MongoDisconnect()
	log.Println("shutting down")
	// log.Println("could take about 15 seconds...")
	// <-ctx.Done()
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	os.Exit(0)
}
