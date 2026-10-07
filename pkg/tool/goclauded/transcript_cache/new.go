package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache/entry"
)

func New(c *claude.Client) *Cache {
	return &Cache{
		Client:   c,
		sessions: map[string]*entry.Session{},
		calls:    map[string]*entry.Call{},
	}
}
