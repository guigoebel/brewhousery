package handlers

import (
	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
	"strconv"
)

func (p *Products) Remove(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for DELETE request")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	p.l.Println("[DEBUG] Retrieved product ID: ", id)
	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Cannot update; Product not found ...", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Cannot delete product ...", http.StatusInternalServerError)
		return
	}
}
