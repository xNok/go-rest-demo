package recipes

import "errors"

type MemStore struct {
	list map[string]Recipe
}

// user represents a recipe
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredient"`
}

type Ingredient struct {
	Name string `json:"name"`
}

func (m MemStore) Add(name string, recipe Recipe) error {
	m.list[name] = recipe
	return nil
}

func (m MemStore) Get(name string) (Recipe, error) {

	if val, ok := m.list["foo"]; ok {
		return val, nil
	}

	return Recipe{}, errors.New("Not found")
}

func (m MemStore) List(name string) (map[string]Recipe, error) {
	return m.list, nil
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
