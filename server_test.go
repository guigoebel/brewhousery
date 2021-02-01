package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)


func TestHelloWorld(t *testing.T) {
    request, error := http.NewRequest("GET", "/", nil)

    if error != nil {
        t.Fatalf("Could not reach endpoint '/' : %v", error)
    }
    // accepts result from HTTP request
    response := httptest.NewRecorder()
    // Initiate HelloWorld function to serve request to recorder
    handlerFunction := http.HandlerFunc(HelloWorld)
    handlerFunction.ServeHTTP(response, request)

    if response.Code != http.StatusOK {
        t.Errorf("Expected response status code 200; got %v", response.Code)
    }

    if response.Body.String() != "Hello, World!" {
        t.Errorf("Expected response message 'Hello, World!'; got %v", response.Body.String())
    }
}


