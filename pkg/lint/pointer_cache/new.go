package pointer_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_entry"
	"hash/maphash"
)

func New() *Cache {
	return &Cache{
		seed:     maphash.MakeSeed(),
		entries:  make(map[string]*pointer_entry.Entry),
		globals:  make(map[string]string),
		Mismatch: func(string, string, string, string) {},
	}
}
