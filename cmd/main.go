package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received new request %s\n", r.URL.Path)
		_, _ = fmt.Fprintf(w, "Hello %s\n", r.URL.Path)
	})

	log.Printf("Server started\n")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
