package data

import (
	"fmt"
	"time"
)

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"         validate:"required"`
	Description    string  `json:"description"`
	Price          float32 `json:"price"        validate:"gt=0"`
	SKU            string  `json:"sku"          validate:"required,sku"`
	ManufacturedOn string  `json:"-"`
	ExpiresOn      string  `json:"-"`
}

var productList = []*Product{
	&Product{
		ID:             1,
		Name:           "Latte",
		Description:    "Frothy milk coffee",
		Price:          2.45,
		SKU:            "COF-LAT-MIL-SUG",
		ManufacturedOn: time.Now().UTC().String(),
		ExpiresOn:      time.Now().UTC().String(),
	},

	&Product{
		ID:             2,
		Name:           "Expresso",
		Description:    "Short and strong coffee without milk",
		Price:          1.99,
		SKU:            "COF-EXP-NOM-SUG",
		ManufacturedOn: time.Now().UTC().String(),
		ExpiresOn:      time.Now().UTC().String(),
	},
}

var ErrProductNotFound = fmt.Errorf("Product not found ...")

type Products []*Product

func GetAllProducts() Products {
	return productList
}

func GetSpecificProduct(id int) (*Product, error) {
	pos := findProductIndex(id)
	if pos == -1 {
		return nil, ErrProductNotFound
	}
	return productList[pos], nil
}

func AddProduct(p *Product) int {
	currentID := productList[len(productList)-1].ID
	p.ID = currentID + 1
	productList = append(productList, p)
	return 1
}

func UpdateProduct(id int, p *Product) error {
	pos := findProductIndex(id)
	if pos == -1 {
		return ErrProductNotFound
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	pos := findProductIndex(id)
	if pos == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

func findProductIndex(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}
