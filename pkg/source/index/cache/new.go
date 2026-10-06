package cache

import "github.com/funtimecoding/soil/pkg/source/index"

func New(
	directory string,
	kinds ...*index.Kind,
) *Cache {
	return &Cache{
		directory: directory,
		kinds:     kinds,
		entries:   make(map[string]*entry),
	}
}
