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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
