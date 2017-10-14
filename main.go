package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, First")
	})

	log.Println("Starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
