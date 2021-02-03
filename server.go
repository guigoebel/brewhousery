package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/saurabmish/Coffee-Shop/handlers"
)

func main() {
	// create logger to inject into reference handler
	l := log.New(os.Stdout, "Coffee shop API service ", log.LstdFlags)
	// create reference to hello handler
	helloHandler := handlers.NewProducts(l)
	// register handler with server
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	// reliability pattern for server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// ensure that service will not be blocked
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// ensure graceful shutdown of server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	l.Println("Received terminate, graceful shutdown", signal)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
