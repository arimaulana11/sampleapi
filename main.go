package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Simple Go API is running 🚀")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello from Go API!"}`)
	})

	log.Println("Server running on port 1234...")
	log.Fatal(http.ListenAndServe("[::]:8300", nil))
}
