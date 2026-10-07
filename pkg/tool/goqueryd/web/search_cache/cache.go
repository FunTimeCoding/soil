package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/search_cache_entry"
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]*search_cache_entry.Entry
	order   []string
	maxSize int
	ttl     time.Duration
}
