package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"time"
)

func (c *Cache) Get(key string) *result.Outcome {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, found := c.entries[key]

	if !found {
		return nil
	}

	if time.Now().After(e.Expiry) {
		delete(c.entries, key)
		c.removeFromOrder(key)

		return nil
	}

	c.promoteInOrder(key)

	return e.Outcome
}
