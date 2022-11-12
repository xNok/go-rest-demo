package main

import (
	"net/http"
	"regexp"
)

var (
	createRecipeRe = regexp.MustCompile(`^\/recipes[\/]*$`)
	listRecipesRe  = regexp.MustCompile(`^\/recipes[\/]*$`)
	getRecipeRe    = regexp.MustCompile(`^\/recipes\/(\d+)$`)
	updateRecipeRe = regexp.MustCompile(`^\/recipes\/(\d+)$`)
	deleteRecipeRe = regexp.MustCompile(`^\/recipes\/(\d+)$`)
)

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
	switch {
	case r.Method == http.MethodPost && createRecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && listRecipesRe.MatchString(r.URL.Path):
		h.ListRecipies(w, r)
		return
	case r.Method == http.MethodPost && getRecipeRe.MatchString(r.URL.Path):
		h.GetRecipie(w, r)
		return
	case r.Method == http.MethodPut && updateRecipeRe.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && deleteRecipeRe.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}
}

func (h *recipes) CreateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipes) ListRecipies(w http.ResponseWriter, r *http.Request) {

}

func (h *recipes) GetRecipie(w http.ResponseWriter, r *http.Request) {

}

func (h *recipes) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipes) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

}

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
