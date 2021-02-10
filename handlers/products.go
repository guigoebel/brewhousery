package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for GET Product")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for POST product")

	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProducts(&product)
}

func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for PUT product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert ID", http.StatusBadRequest)
	}

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		product := data.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Error reading product ...", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
