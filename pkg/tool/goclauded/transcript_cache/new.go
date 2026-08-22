package transcript_cache

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude"

func New(c *claude.Client) *Cache {
	return &Cache{
		Client:   c,
		sessions: map[string]*sessionEntry{},
		calls:    map[string]*callEntry{},
	}
}
