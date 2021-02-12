package handlers

import (
	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"net/http"
	"strconv"
)

func (p *Products) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for GET all products")
	listProducts := data.GetAll()
	err := listProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}

func (p *Products) GetSpecificProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint for GET one product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	p.l.Println("[DEBUG] Product ID is: ", id)
	product, err := data.GetSpecific(id)
	if err != nil {
		http.Error(w, "Unable to find product with the given ID ...", http.StatusInternalServerError)
	}
	p.l.Println("[DEBUG] Retrieved product data successfully!")
	val := product.ToJSONSingle(w)
	if val != nil {
		http.Error(w, "Unable to parse values to JSON", http.StatusInternalServerError)
	}
}
