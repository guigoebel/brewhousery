package handlers

import (
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p Products) Remove(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for DELETE request")
	w.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	p.l.Println("[DEBUG] Retrieved product ID: ", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Couldn't delete product; ID not found ..."}, w)
		p.l.Println("[ERROR] Fetching product with given ID ...", err)
		return

	}
	if err != nil {
		http.Error(w, "Cannot delete product ...", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
