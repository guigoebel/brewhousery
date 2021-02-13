package handlers

import (
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
)

func (p Products) Modify(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for PUT request")

	id := getProductID(r)
	product := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Retrieved product from data store")

	err := data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Cannot update; Product not found ...", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Cannot update product ...", http.StatusInternalServerError)
		return
	}
}
