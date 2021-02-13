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
