package handlers

import (
	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
	"strconv"
)

func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for PUT product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
