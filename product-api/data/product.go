package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

type Products []*Product

func (products *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(products)
}

func (product *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

func GetProducts() Products {
	return productList
}

func AddProduct(product *Product) {
	product.ID = getNextId()
	productList = append(productList, product)
}

func UpdateProduct(id int, product *Product) error {
	_, position, err := findProduct(id)

	if err != nil {
		return err
	}

	product.ID = id
	productList[position] = product

	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextId() int {
	product := productList[len(productList)-1]
	return product.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Nescau",
		Description: "Cereal",
		Price:       3.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Coffe",
		Description: "Black coffee powder",
		Price:       4.45,
		SKU:         "xah232",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
