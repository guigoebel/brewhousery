package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saurabmish/Coffee-Shop/data"
)

// Products is a struct which implements the HTTP handler interface
type Products struct {
	l *log.Logger
}

// NewProducts is going to accept a logger and return products as a reference
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Hit endpoint '/'")
	fmt.Fprintf(w, "Welcome to coffee shop!")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}
