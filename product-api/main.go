package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "products-api", log.LstdFlags)

	productHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port:", server.Addr)
		err := server.ListenAndServe()

		if err != nil {
			logger.Printf("Error starting server %s \n", err)
			os.Exit(1)
		}

	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	logger.Println("Interrupt signal received, gracefully shutdown", sig)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(timeoutContext)

}
