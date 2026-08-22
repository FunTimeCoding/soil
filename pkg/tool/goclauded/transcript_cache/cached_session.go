package transcript_cache

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func (c *Cache) cachedSession(identifier string) *session.Session {
	modTime, size, exists := c.fileState(identifier)

	if !exists {
		return session.Stub()
	}

	c.mutex.RLock()
	entry, found := c.sessions[identifier]
	c.mutex.RUnlock()

	if found && entry.ModTime.Equal(modTime) && entry.Size == size {
		return entry.Session
	}

	s := c.Session(identifier)
	c.mutex.Lock()
	c.sessions[identifier] = &sessionEntry{
		ModTime: modTime,
		Size:    size,
		Session: s,
	}
	c.mutex.Unlock()

	return s
}
