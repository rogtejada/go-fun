package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorilla-product-api/data"
	"gorilla-product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "products-api", log.LstdFlags)
	validation := data.NewValidation()

	productHandler := handlers.NewProducts(logger, validation)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.GetById)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS - allowing all origins
	corsHandler := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(serveMux),
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
