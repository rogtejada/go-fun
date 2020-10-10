package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorilla-product-api/data"
	"net/http"
	"strconv"
)



// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// UpdateProduct handles PUT requests to update products
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
		p.logger.Println("[ERROR]", err)
		http.Error(writer, fmt.Sprintf("Error updating product: %s", err), http.StatusInternalServerError)
		return
	}
}
