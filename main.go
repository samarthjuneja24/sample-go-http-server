package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sample-go-http-server/routes"
	"syscall"
	"time"
)

func main() {
	router := routes.SetupRouter()
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("Starting server on port 8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Setting up graceful shutdown.
	gracefulShutdown(server)
}

// gracefulShutdown handles graceful shutdown of the HTTP server.
func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server gracefully shutdown")
}
