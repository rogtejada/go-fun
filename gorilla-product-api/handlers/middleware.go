package handlers

import (
	"context"
	"fmt"
	"gorilla-product-api/data"
	"net/http"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		product := data.Product{}

		err := data.FromJSON(&product, request.Body)

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

