package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/search_cache_entry"
	"time"
)

func New(maxSize int) *Cache {
	return &Cache{
		entries: map[string]*search_cache_entry.Entry{},
		maxSize: maxSize,
		ttl:     15 * time.Minute,
	}
}
