package recipes

// represents a recipe
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredient"`
}

// represents individual ingredients
type Ingredient struct {
	Name string `json:"name"`
}
