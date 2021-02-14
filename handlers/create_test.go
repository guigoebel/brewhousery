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

func CreateRouter() *mux.Router {
	l := log.New(os.Stdout, "[TEST] Coffee shop API service ", log.LstdFlags)
	v := data.NewValidation()

	coffeeHandler := NewProducts(l, v)
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/coffee/add", coffeeHandler.Add).Methods("POST")
	serveMux.Use(coffeeHandler.MiddlewareProductValidation)
	return serveMux
}

func TestCreateProductWithoutID(t *testing.T) {
	payload := []byte(`{"name": "Tea", "description": "A nice cup of tea!", "price": 2.0, "sku": "TEA-WAT-MIL-SUG"}`)
	request, _ := http.NewRequest("POST", "/coffee/add", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code, "Product created successfully!") // http.StatusCreated = 201
	assert.Equal(t, `{"message":"Product created successfully!"}`, strings.TrimRight(response.Body.String(), "\n"))
}

func TestCreateProductWithID(t *testing.T) {
	payload := []byte(`{"id": 12, "name": "Caramel Macchiato", "description": "Freshly steamed milk with vanilla-flavored syrup marked with espresso and topped with caramel drizzle.", "price": 4.0, "sku": "CAR-MAC-VAN-ESP"}`)
	request, _ := http.NewRequest("POST", "/coffee/add", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	CreateRouter().ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code, "Product ID is auto-incremented, specifying it doesn't have any effect")
	assert.Equal(t, `{"message":"Product created successfully!"}`, strings.TrimRight(response.Body.String(), "\n"))
}
