package handlers

import (
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
)

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for POST product")

	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProducts(&product)
}
