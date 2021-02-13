package handlers

import (
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
)

func (p *Products) Remove(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for DELETE request")

	id := getProductID(r)
	p.l.Println("[DEBUG] Retrieved product ID: ", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Cannot update; Product not found ...", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Cannot delete product ...", http.StatusInternalServerError)
		return
	}
}
