package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float32 `json:"price"`
	SKU            string  `json:"sku"`
	ManufacturedOn string  `json:"-"`
	ExpiresOn      string  `json:"expiry"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
