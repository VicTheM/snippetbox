package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Registering routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on port :4001...")
	log.Println("Server is running!")

	err := http.ListenAndServe(":4001", mux)

	log.Fatal(err)
}
