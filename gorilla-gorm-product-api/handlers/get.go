package handlers

import (
	"gorilla-product-api/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse
func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-type", "application/json")
	listOfProducts := data.GetProducts()
	err := data.ToJSON(listOfProducts, writer)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse
func (p *Products) GetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-type", "application/json")
	id := getProductID(request)

	p.logger.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.logger.Println("[ERROR] fetching product", err)
		writer.WriteHeader(http.StatusNotFound)
		return
	default:
		p.logger.Println("[ERROR] fetching product", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(prod, writer)
	if err != nil {
		p.logger.Println("[ERROR] serializing product", err)
	}
}