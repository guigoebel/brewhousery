package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidProduct(t *testing.T) {
	p := &Product{
		Name:        "Cappuccino",
		Description: "Espresso-based coffee prepared with steamed milk foam",
		Price:       3.50,
		SKU:         "COF-CAP-MIL-FOA",
	}

	v := NewValidation()
	err := v.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}

func TestProductWithoutName(t *testing.T) {
	p := Product{
		Description: "No description present ...",
		Price:       4.50,
		SKU:         "ABC-IJK-PQR-XYZ",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestInvalidSKU(t *testing.T) {
	p := &Product{
		Name:        "Nitro Cold Brew",
		Description: "Nitrogen infused cold brew with super-smooth taste",
		Price:       5.50,
		SKU:         "NIT-COL-BRE",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestInvalidPrice(t *testing.T) {
	p := &Product{
		Name:        "Eggnog Latte",
		Description: "Nitrogen infused cold brew with super-smooth taste",
		Price:       0,
		SKU:         "EGG-NOG-LAT-NUT",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}
