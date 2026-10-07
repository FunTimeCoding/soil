package cache

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/types/cache_entry"
)

func New(
	directory string,
	kinds ...*index.Kind,
) *Cache {
	return &Cache{
		directory: directory,
		kinds:     kinds,
		entries:   make(map[string]*cache_entry.Entry),
	}
}
