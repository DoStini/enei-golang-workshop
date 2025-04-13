package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Enei!\n"))
	})
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
