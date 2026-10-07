package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache/entry"
)

func (c *Cache) cachedSession(identifier string) *session.Session {
	modTime, size, exists := c.fileState(identifier)

	if !exists {
		return session.Stub()
	}

	c.mutex.RLock()
	e, found := c.sessions[identifier]
	c.mutex.RUnlock()

	if found && e.ModTime.Equal(modTime) && e.Size == size {
		return e.Session
	}

	s := c.Session(identifier)
	c.mutex.Lock()
	c.sessions[identifier] = entry.NewSession(modTime, size, s)
	c.mutex.Unlock()

	return s
}
