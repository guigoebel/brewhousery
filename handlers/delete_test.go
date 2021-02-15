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

func DeleteRouter() *mux.Router {
	l := log.New(os.Stdout, "[TEST] Coffee shop API service ", log.LstdFlags)
	v := data.NewValidation()

	coffeeHandler := NewProducts(l, v)
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/coffee/remove/{id:[0-9]+}", coffeeHandler.Remove).Methods("DELETE")
	return serveMux
}

func TestDeleteExistingProduct(t *testing.T) {

	request, _ := http.NewRequest("DELETE", "/coffee/remove/1", nil)
	response := httptest.NewRecorder()
	DeleteRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code, "Product deleted successfully!")
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestDeleteNonExistingProduct(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/coffee/remove/19", nil)
	response := httptest.NewRecorder()
	DeleteRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code, "Product not found ...") // http.StatusNotFound = 404
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Equal(t, `{"message":"Couldn't delete product; ID not found ..."}`, strings.TrimRight(response.Body.String(), "\n"))
}
