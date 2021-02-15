package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"github.com/stretchr/testify/assert"
)

func MiddlewareRouter() *mux.Router {
	l := log.New(os.Stdout, "[TEST] Coffee shop API service ", log.LstdFlags)
	v := data.NewValidation()

	coffeeHandler := NewProducts(l, v)
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/coffee/add", coffeeHandler.Add).Methods("POST")
	serveMux.HandleFunc("/coffee/modify/{id:[0-9]+}", coffeeHandler.Modify).Methods("PUT")
	serveMux.Use(coffeeHandler.MiddlewareProductValidation)
	return serveMux
}

func TestProductWithInsufficientFieldsUsingUpdate(t *testing.T) {
	payload := []byte(`{"name": "Mocha", "sku": "COF-MOC-VAR-LAT"}`)
	request, _ := http.NewRequest("PUT", "/coffee/modify/1", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	MiddlewareRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code) // http.StatusUnprocessableEntity = 422
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestProductWithWrongValueTypeUsingUpdate(t *testing.T) {
	payload := []byte(`{"name": 1234, "description": "Chocolate-flavoured variant of latte", "sku": "COF-MOC-VAR-LAT"}`)
	request, _ := http.NewRequest("PUT", "/coffee/modify/1", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	MiddlewareRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code) // http.StatusBadRequest = 400
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestProductWithExtraFieldsUsingCreate(t *testing.T) {
	payload := []byte(`{"name": "Caramel Macchiato", "blob1": 12, "description": "Freshly steamed milk with vanilla-flavored syrup marked with espresso and topped with caramel drizzle.", "price": 4.0, "sku": "CAR-MAC-VAN-ESP", "blob2": 2}`)
	request, _ := http.NewRequest("POST", "/coffee/add", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	MiddlewareRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code, "Extra fields will be discarded")
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func TestProductWithWrongValueTypeUsingCreate(t *testing.T) {
	payload := []byte(`{"name": "Caramel Macchiato", "blob1": 12, "description": None, "price": 4.0, "sku": "CAR-MAC-VAN-ESP", "blob2": 2}`)
	request, _ := http.NewRequest("PUT", "/coffee/modify/1", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	MiddlewareRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}
