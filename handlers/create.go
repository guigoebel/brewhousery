package handlers

import (
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p Products) Add(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for POST request")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Adding product to list")
	err := data.AddProduct(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: "Couldn't create product; ID not found ..."}, w)
		p.l.Println("[ERROR] Unable to parse values to JSON ...", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	data.ToJSON(&GenericError{Message: "Product created successfully!"}, w)
}
