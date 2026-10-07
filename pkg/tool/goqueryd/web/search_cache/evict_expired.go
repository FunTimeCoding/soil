package search_cache

import "time"

func (c *Cache) evictExpired() {
	now := time.Now()
	var kept []string

	for _, key := range c.order {
		if e, found := c.entries[key]; found && now.Before(e.Expiry) {
			kept = append(kept, key)
		} else {
			delete(c.entries, key)
		}
	}

	c.order = kept
}
