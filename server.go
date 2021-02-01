package main

import (
	"fmt"
	"log"
	"net/http"
)


func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}


func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
