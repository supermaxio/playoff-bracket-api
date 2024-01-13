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
	"github.com/supermaxio/nflplayoffbracket/service"
)

func main() {
	log.Println("starting up")
	// Setup configurations
	config.Setup()

	// Ping db
	database.MongoConnect()

	// Setup a ticker that triggers every minute
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()

    // Context for graceful shutdown
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

	// Goroutine for refreshing scores
	go func() {
		for {
			select {
			case <-ticker.C:
				// Call your function here
				_, err := service.RefreshScores()
				if err != nil {
					fmt.Println("Error refreshing scores:", err)
				} else {
					fmt.Println("refreshed the scores!")
				}
			case <-ctx.Done():
				// Context was cancelled, exit the goroutine
                fmt.Println("Refresh scores goroutine exiting...")
				return
			}
		}
	}()

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

	sigint := make(chan os.Signal, 1)
	sigterm := make(chan os.Signal, 1)
	sigkill := make(chan os.Signal, 1)
	sigquit := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigkill, syscall.SIGKILL)
	signal.Notify(sigquit, syscall.SIGQUIT)

	// Wait for the SIGTERM signal or the context to expire
	select {
	case <-sigint:
		log.Println("Received SIGINT, shutting down...")
	case <-sigterm:
		log.Println("Received SIGTERM, shutting down...")
	case <-sigkill:
		log.Println("Received SIGKILL, shutting down...")
	case <-sigquit:
		log.Println("Received SIGQUIT, shutting down...")
	}

	log.Println("Signal the goroutine to stop")
    // Cancel the context, which should unblock ctx.Done()
	cancel() 

	log.Println("Create a deadline to wait for.")

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// go srv.Shutdown(ctx)
	server.Shutdown(ctx)

	database.MongoDisconnect()
	log.Println("shutting down")
	// log.Println("could take about 15 seconds...")
	// <-ctx.Done()
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	os.Exit(0)
}
