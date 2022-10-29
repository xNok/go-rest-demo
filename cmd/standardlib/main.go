package main

import "net/http"

func main() {

	// Create an HTTP server
	mux := http.NewServeMux()

	// Tegister the / (home) route
	mux.Handle("/", &home{})
}

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
