package handlers

import (
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

func GetRouter() *mux.Router {
	l := log.New(os.Stdout, "[TEST] Coffee shop API service ", log.LstdFlags)
	v := data.NewValidation()

	coffeeHandler := NewProducts(l, v)
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/coffee/get/all", coffeeHandler.RetrieveAll).Methods("GET")
	serveMux.HandleFunc("/coffee/get/{id:[0-9]+}", coffeeHandler.RetrieveSingle).Methods("GET")
	return serveMux
}

func TestRetrieveAllProducts(t *testing.T) {
	request, _ := http.NewRequest("GET", "/coffee/get/all", nil)
	response := httptest.NewRecorder()
	GetRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestRetrieveSingleProduct(t *testing.T) {
	request, _ := http.NewRequest("GET", "/coffee/get/1", nil)
	response := httptest.NewRecorder()
	GetRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Product Found!")
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestRetrieveNonExistingProduct(t *testing.T) {
	request, _ := http.NewRequest("GET", "/coffee/get/20", nil)
	response := httptest.NewRecorder()
	GetRouter().ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "Unable to find product with the given ID ...")
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Equal(t, `{"message":"Product not found ..."}`, strings.TrimRight(response.Body.String(), "\n"))
}
