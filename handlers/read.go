package handlers

import (
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
)

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for GET Product")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}
