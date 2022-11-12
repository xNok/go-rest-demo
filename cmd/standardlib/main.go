package main

import "net/http"

func main() {

	// Create a new request multiplexer
	// Takes incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &home{})
	mux.Handle("/recipes", &recipes{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

type recipes struct{}

func (h *recipes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my recipe page"))
}

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
