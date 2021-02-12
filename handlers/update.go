package handlers

import (
	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
	"strconv"
)

func (p Products) Update(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for PUT request")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Cannot update; Product not found ...", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Cannot update product ...", http.StatusInternalServerError)
		return
	}
}
