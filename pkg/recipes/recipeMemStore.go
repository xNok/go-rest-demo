package recipes

import "errors"

var (
	NotFoundErr = errors.New("not found")
)

type MemStore struct {
	list map[string]Recipe
}

func NewMemStore() *MemStore {
	return &MemStore{}
}

func (m MemStore) Add(name string, recipe Recipe) error {
	m.list[name] = recipe
	return nil
}

func (m MemStore) Get(name string) (Recipe, error) {

	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return Recipe{}, NotFoundErr
}

func (m MemStore) List() (map[string]Recipe, error) {
	return m.list, nil
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
