package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/handlers"
)

func main() {
	l := log.New(os.Stdout, "Coffee shop API service ", log.LstdFlags)
	coffeeHandler := handlers.NewProducts(l)
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", coffeeHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", coffeeHandler.UpdateProduct)
	putRouter.Use(coffeeHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", coffeeHandler.AddProducts)
	postRouter.Use(coffeeHandler.MiddlewareProductValidation)

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
