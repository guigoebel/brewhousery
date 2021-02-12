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
	ExpiresOn      string  `json:"expiry"`
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

func GetProducts() Products {
	return productList
}

func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}
