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

func TestFindProductIndex(t *testing.T) {
	got := findProductIndex(1)
	if got == -1 {
		t.Errorf("Could not find product with given ID")
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
