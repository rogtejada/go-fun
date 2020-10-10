// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /products
// Version: 1.0.0
//
// Consumers:
// 	- application/json
//
// Produces:
// 	- application/json
// swagger:meta

package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorilla-product-api/data"
	"log"
	"net/http"
	"strconv"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

type Products struct {
	logger *log.Logger
	validator *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}