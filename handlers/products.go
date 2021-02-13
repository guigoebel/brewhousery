package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type KeyProduct struct{}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// getProductID converts string ID from the request URI to integer ID
// Used in handler functions 'RetrieveSingle', 'Update', and 'Remove'
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err == nil {
		return id
	}
	panic(err)
}
