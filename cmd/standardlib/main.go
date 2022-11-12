package main

import (
	"net/http"
	"regexp"

	"github.com/xNok/go-rest-demo/recipes"
)

var (
	RecipeRe       = regexp.MustCompile(`^\/recipes[\/]*$`)
	RecipeReWithID = regexp.MustCompile(`^\/recipes\/(\d+)$`)
)

func main() {

	// Create a new request multiplexer
	// Takes incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &recipesHandler{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

type recipesHandler struct{}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}

func (h *recipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipies(w, r)
		return
	case r.Method == http.MethodPost && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipie(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}
}

func (h *recipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipesHandler) ListRecipies(w http.ResponseWriter, r *http.Request) {

}

func (h *recipesHandler) GetRecipie(w http.ResponseWriter, r *http.Request) {

}

func (h *recipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
