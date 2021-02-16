package handlers

import (
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p Products) Modify(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for PUT request")
	w.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	p.l.Println("[DEBUG] Retrieved product ID from URL: ", id)
	product := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Retrieved product from data store")

	err := data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Couldn't update product; ID not found ..."}, w)
		p.l.Println("[ERROR] Fetching product with given ID ...", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
