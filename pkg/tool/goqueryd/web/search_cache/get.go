package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"time"
)

func (c *Cache) Get(key string) *store.SearchOutcome {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, found := c.entries[key]

	if !found {
		return nil
	}

	if time.Now().After(e.expiry) {
		delete(c.entries, key)
		c.removeFromOrder(key)

		return nil
	}

	c.promoteInOrder(key)

	return e.outcome
}
