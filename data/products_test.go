package data

import (
	"testing"
	"time"
)

func TestGetAllProducts(t *testing.T) {
	got := GetAllProducts()
	if got == nil {
		t.Errorf("No values in data store ...")
	}
}

func TestGetSpecificExistingProduct(t *testing.T) {
	prod, err := GetSpecificProduct(1)
	if prod == nil && err == ErrProductNotFound {
		t.Errorf("Could not get product with given ID ...")
	}
}

func TestGetSpecificNonExistingProduct(t *testing.T) {
	prod, err := GetSpecificProduct(12)
	if prod != nil && err == ErrProductNotFound {
		t.Errorf("Could not get product with given ID as it doesn't exist!")
	}
}

func TestFindExistingProductIndex(t *testing.T) {
	got := findProductIndex(1)
	if got == -1 {
		t.Errorf("Could not find existing product ...")
	}
}

func TestFindNonExistingProductIndex(t *testing.T) {
	got := findProductIndex(10)
	if got != -1 {
		t.Errorf("Could not find product as it doesn't exist!")
	}
}

func TestDeleteExistingProduct(t *testing.T) {
	got := DeleteProduct(2)
	if got != nil {
		t.Errorf("Could not delete existing product ...")
	}
}

func TestDeleteNonExistingProduct(t *testing.T) {
	got := DeleteProduct(20)
	if got != ErrProductNotFound {
		t.Errorf("Could not delete product as it doesn't exist!")
	}
}

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

	got := AddProduct(payload)
	want := 1
	if got != want {
		t.Errorf("Cannot add product ...")
	}
}

func TestUpdateExistingProduct(t *testing.T) {
	payload := &Product{
		12,
		"Honey Oatmilk Latte",
		"Oatmilk and espresso roast combined with a hint of honey",
		3.50,
		"HON-OAT-LAT-ESP",
		time.Now().UTC().String(),
		time.Now().UTC().String(),
	}

	got := UpdateProduct(1, payload)
	if got == ErrProductNotFound {
		t.Errorf("Could not update product with given ID ...")
	}
}

func TestUpdateNonExistingProduct(t *testing.T) {
	payload := &Product{
		12,
		"Honey Oatmilk Latte",
		"Oatmilk and espresso roast combined with a hint of honey",
		3.50,
		"HON-OAT-LAT-ESP",
		time.Now().UTC().String(),
		time.Now().UTC().String(),
	}

	got := UpdateProduct(12, payload)
	if got == nil {
		t.Errorf("Could not update product with given ID as it doesn't exist!")
	}
}
