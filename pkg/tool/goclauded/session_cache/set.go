package session_cache

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"

func (c *Cache) Set(
	identifier string,
	state *tracker.State,
) {
	c.mutex.Lock()
	c.entries[identifier] = state
	c.mutex.Unlock()
}
