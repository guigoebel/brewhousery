package handlers

import (
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

func (p *Products) RetrieveAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for READ request")
	w.Header().Add("Content-Type", "application/json")

	products := data.GetAllProducts()
	p.l.Println("[DEBUG] Retrieved all products")

	err := data.ToJSON(products, w)
	if err != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}

func (p *Products) RetrieveSingle(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for READ request")
	w.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	p.l.Println("[DEBUG] Retrieved product ID: ", id)

	product, err := data.GetSpecificProduct(id)
	if err != nil {
		http.Error(w, "Unable to find product with the given ID ...", http.StatusNotFound)
		return
	}

	val := data.ToJSON(product, w)
	if val != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
		return
	}
}
