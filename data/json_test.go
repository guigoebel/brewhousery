package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductNameToJSON(t *testing.T) {
	p := []*Product{
		&Product{
			Name: "Flat White",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(p, b)
	assert.NoError(t, err)
}

func TestProductNameAndPriceToJSON(t *testing.T) {
	p := []*Product{
		&Product{
			Name:  "Flat White",
			Price: 2.50,
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(p, b)
	assert.NoError(t, err)
}

func TestProductNamePriceAndSKUToJSON(t *testing.T) {
	p := []*Product{
		&Product{
			Name:  "Flat White",
			Price: 2.50,
			SKU:   "FLA-WHI-ESP-MIL",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(p, b)
	assert.NoError(t, err)
}

func TestProductDetailsToJSON(t *testing.T) {
	p := []*Product{
		&Product{
			Name:        "Flat White",
			Description: "Espresso-based drink with steamed milk.",
			Price:       2.50,
			SKU:         "FLA-WHI-ESP-MIL",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(p, b)
	assert.NoError(t, err)
}

func TestProductNameFromJSON(t *testing.T) {
	payload := `{"name": "Flat White"}`
	b := bytes.NewBufferString(payload)
	err := FromJSON(&Product{}, b)
	assert.NoError(t, err)
}

func TestProductNameAndPriceFromJSON(t *testing.T) {
	payload := `{"name": "Flat White", "price": 2.50}`
	b := bytes.NewBufferString(payload)
	err := FromJSON(&Product{}, b)
	assert.NoError(t, err)
}

func TestProductNamePriceAndSKUFromJSON(t *testing.T) {
	payload := `{"name": "Flat White", "price": 2.50, "sku": "FLA-WHI-ESP-MIL"}`
	b := bytes.NewBufferString(payload)
	err := FromJSON(&Product{}, b)
	assert.NoError(t, err)
}

func TestProductDetailsFromJSON(t *testing.T) {
	payload := `{"name": "Flat White", "description":"Espresso-based drink with steamed milk.", "price": 2.50, "sku": "FLA-WHI-ESP-MIL"}`
	b := bytes.NewBufferString(payload)
	err := FromJSON(&Product{}, b)
	assert.NoError(t, err)
}
