package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"gorilla-product-api/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(writer)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {

	product := request.Context().Value(KeyProduct{}).(data.Product)

	p.logger.Printf("Product: %#v", product)
	data.AddProduct(&product)
}

func (p *Products) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		http.Error(writer, "Unable to parse id", http.StatusBadRequest)
		return
	}

	product := request.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)

	if err == data.ErrProductNotFound {
		http.Error(writer, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		product := data.Product{}

		err := product.FromJSON(request.Body)

		if err != nil {
			p.logger.Println("[ERROR]", err)
			http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		err = product.Validate()

		if err != nil {
			p.logger.Println("[ERROR] validating product", err)
			http.Error(writer, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, product)
		req := request.WithContext(ctx)

		next.ServeHTTP(writer, req)
	})
}
