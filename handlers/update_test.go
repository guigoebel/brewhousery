package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"github.com/stretchr/testify/assert"
)

func PutRouter() *mux.Router {
	l := log.New(os.Stdout, "[TEST] Coffee shop API service ", log.LstdFlags)
	v := data.NewValidation()

	coffeeHandler := NewProducts(l, v)
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/coffee/modify/{id:[0-9]+}", coffeeHandler.Modify).Methods("PUT")
	serveMux.Use(coffeeHandler.MiddlewareProductValidation)
	return serveMux
}

func TestUpdateExistingProduct(t *testing.T) {
	payload := []byte(`{"name": "Mocha", "description": "Chocolate-flavoured variant of latte", "price": 3.00, "sku": "COF-MOC-VAR-LAT"}`)
	request, _ := http.NewRequest("PUT", "/coffee/modify/2", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	PutRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code, "Product updated successfully!") // http.StatusNoContent = 204
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestUpdateNonExistingProduct(t *testing.T) {
	payload := []byte(`{"name": "Mocha", "description": "Chocolate-flavoured variant of latte", "price": 3.00, "sku": "COF-MOC-VAR-LAT"}`)
	request, _ := http.NewRequest("PUT", "/coffee/modify/19", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	PutRouter().ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "Product not found ...")
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Equal(t, `{"message":"Couldn't update product; ID not found ..."}`, strings.TrimRight(response.Body.String(), "\n"))
}
