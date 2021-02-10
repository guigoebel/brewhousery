package data

import "testing"

func TestSKUValidation(t *testing.T) {
	p := &Product{
		Name:        "Cappuccino",
		Description: "Espresso-based coffee prepared with steamed milk foam",
		Price:       3.50,
		SKU:         "COF-CAP-MIL-FOA",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
