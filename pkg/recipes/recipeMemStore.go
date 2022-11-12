package recipes

import "errors"

type MemStore struct {
	list map[string]Recipe
}

func (m MemStore) Add(name string, recipe Recipe) error {
	m.list[name] = recipe
	return nil
}

func (m MemStore) Get(name string) (Recipe, error) {

	if val, ok := m.list["foo"]; ok {
		return val, nil
	}

	return Recipe{}, errors.New("not found")
}

func (m MemStore) List() (map[string]Recipe, error) {
	return m.list, nil
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
