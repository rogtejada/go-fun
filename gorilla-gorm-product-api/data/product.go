package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


var db *gorm.DB

const username = "postgres"
const password = "password"
const dbName = "test"
const dbHost = "localhost"

var ErrProductNotFound = fmt.Errorf("product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	gorm.Model
	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}
type Products []*Product

func (product *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(product)
}

func GetProducts() []Product {
	var listOfProducts []Product
	db.Find(&listOfProducts)
	return listOfProducts
}

func GetProductByID(id int) (Product, error) {
	var product Product
	err := db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return product, ErrProductNotFound
	}
	return product, nil
}

func AddProduct(product *Product) *Product {
	db.Create(product)
	return product
}

func UpdateProduct(id int, product *Product) (Product, error) {
	var productToUpdate Product
	err := db.Where("id = ?", id).Find(&productToUpdate).Error

	if err != nil {
		return productToUpdate, ErrProductNotFound
	}

	productToUpdate.Name = product.Name
	productToUpdate.Price = product.Price
	productToUpdate.Description = product.Description
	productToUpdate.SKU = product.SKU

	err = db.Save(productToUpdate).Error

	if err != nil {
		return productToUpdate, err
	}

	return productToUpdate, nil
}

func DeleteProduct(id int) {
	var product Product
	db.Where("id = ?", id).Find(&product)
	db.Delete(&product)
}

func InitDatabase() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Product{})
}