package session_cache

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"

func (c *Cache) Get(identifier string) (*tracker.State, bool) {
	c.mutex.Lock()
	state, exists := c.entries[identifier]
	c.mutex.Unlock()

	return state, exists
}
