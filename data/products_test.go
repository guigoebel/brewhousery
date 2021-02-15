package data

import (
	"testing"
)

func TestGetAllProducts(t *testing.T) {
	got := GetAllProducts()
	if got == nil {
		t.Errorf("No values in data store ...")
	}
}

func TestGetSpecificProduct(t *testing.T) {
	prod, err := GetSpecificProduct(1)
	if prod == nil && err == ErrProductNotFound {
		t.Errorf("Could not find product with given ID")
	}
}

func TestFindExistingProductIndex(t *testing.T) {
	got := findProductIndex(1)
	if got == -1 {
		t.Errorf("Could not find product with given ID")
	}
}

func TestFindNonExistingProductIndex(t *testing.T) {
	got := findProductIndex(10)
	if got != -1 {
		t.Errorf("Product doesn't exist, cannot retrieve ...")
	}
}

func TestDeleteExistingProduct(t *testing.T) {
	got := DeleteProduct(2)
	if got != nil {
		t.Errorf("Could not delete product with given ID")
	}
}

func TestDeleteNonExistingProduct(t *testing.T) {
	got := DeleteProduct(20)
	if got != ErrProductNotFound {
		t.Errorf("Product doesn't exist, cannot delete ...")
	}
}

/*
func TestAddProduct(t *testing.T) {
	payload := &Product{
		12,
		"Flat White",
		"Espresso-based drink with steamed milk.",
		2.50,
		"FLA-WHI-ESP-MIL",
		time.Now().UTC().String(),
		time.Now().UTC().String(),
	}
	_, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("Cannot add product ...")
	}

	if AddProduct(payload) != nil {
		t.Errorf("Cannot add product ...")
	}
}
*/
