package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
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

func (product *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(product)
}

func validateSKU(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	mathches := regex.FindAllString(fl.Field().String(), -1)

	if len(mathches) != 1 {
		return false
	}

	return true
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
		SKU:         "abc-xas-das",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Coffe",
		Description: "Black coffee powder",
		Price:       4.45,
		SKU:         "xah-asd-das",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
