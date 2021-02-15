package handlers

import (
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p Products) RetrieveAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for READ request")
	w.Header().Add("Content-Type", "application/json")

	products := data.GetAllProducts()
	p.l.Println("[DEBUG] Retrieved all products")

	data.ToJSON(products, w)
}

func (p Products) RetrieveSingle(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for READ request")
	w.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	p.l.Println("[DEBUG] Retrieved product ID from URL: ", id)

	_, err := data.GetSpecificProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found ..."}, w)
		p.l.Println("[ERROR] Fetching product with given ID ...", err)
		return
	}
}
