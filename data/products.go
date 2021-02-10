package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
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

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[A-Z]+-[A-Z]+-[A-Z]+-[A-Z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

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

var ErrProductNotFound = fmt.Errorf("Product not found ...")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
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
