package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Hello is a struct which implements interface HTTP handler
type Hello struct {
	l *log.Logger
}

// NewHello is going to accept a logger and return Hello handler as a reference
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// method satisfies the HTTP handler interface
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hit endpoint '/'")
	fmt.Fprintf(w, "Hello, World!")
}
