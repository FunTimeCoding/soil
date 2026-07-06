package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"time"
)

func (c *Cache) Put(
	key string,
	outcome *store.SearchOutcome,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.evictExpired()

	if len(c.entries) >= c.maxSize {
		oldest := c.order[0]
		delete(c.entries, oldest)
		c.order = c.order[1:]
	}

	c.entries[key] = &entry{
		outcome: outcome,
		expiry:  time.Now().Add(c.ttl),
	}
	c.removeFromOrder(key)
	c.order = append(c.order, key)
}
