package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
)

type KeyProduct struct{}

type Products struct {
	l *log.Logger
	v *data.Validation
}

type GenericError struct {
	Message string `json:"message"`
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// getProductID converts string ID from the request URI to integer ID
// Used in handler functions 'RetrieveSingle', 'Update', and 'Remove'
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err == nil {
		return id
	}
	return -1
}
