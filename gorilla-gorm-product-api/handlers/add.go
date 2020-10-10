package handlers

import (
	"gorilla-product-api/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddProduct handles POST requests to add new products
func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-type", "application/json")
	product := request.Context().Value(KeyProduct{}).(data.Product)
	savedProduct := data.AddProduct(&product)

	err := data.ToJSON(savedProduct, writer)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
	}
}
