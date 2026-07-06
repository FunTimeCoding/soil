package session_cache

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"

func (c *Cache) GetOrCreate(identifier string) *tracker.State {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	state, exists := c.entries[identifier]

	if !exists {
		state = tracker.New()
		c.entries[identifier] = state
	}

	return state
}
