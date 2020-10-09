package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	p.logger.Println("Request received", request.Method, request.URL)

	if request.Method == http.MethodGet {
		p.getProducts(responseWriter, request)
		return
	}

	if request.Method == http.MethodPost {
		p.addProduct(responseWriter, request)
		return
	}

	if request.Method == http.MethodPut {
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(request.URL.Path, -1)

		if len(group) != 1 || len(group[0]) != 2 {
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(group[0][1])

		if err != nil {
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, responseWriter, request)
		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(writer)

	if err != nil {
		http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request) {

	product := &data.Product{}
	err := product.FromJSON(request.Body)

	if err != nil {
		http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.logger.Printf("Product: %#v", product)

	data.AddProduct(product)
}

func (p *Products) updateProduct(id int, writer http.ResponseWriter, request *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(request.Body)

	if err != nil {
		http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(writer, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}
}
