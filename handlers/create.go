package handlers

import (
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
)

func (p *Products) Add(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Endpoint for POST request")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Adding product to list")
	data.AddProduct(&product)
}
