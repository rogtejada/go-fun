package handlers

import (
	"gorilla-product-api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.logger.Println("[DEBUG] deleting record id", id)

	data.DeleteProduct(id)

	rw.WriteHeader(http.StatusNoContent)
}
