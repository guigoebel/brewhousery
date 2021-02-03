package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCoffeeShop(t *testing.T) {

	l := log.New(os.Stdout, "Test coffee shop API service ", log.LstdFlags)
	helloHandler := NewProducts(l)

	request, error := http.NewRequest("GET", "/", nil)

	if error != nil {
		t.Fatalf("Could not reach endpoint '/' : %v", error)
	}

	response := httptest.NewRecorder()

	handlerFunction := http.HandlerFunc(helloHandler.ServeHTTP)
	handlerFunction.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected response status code 200; got %v", response.Code)
	}
}
