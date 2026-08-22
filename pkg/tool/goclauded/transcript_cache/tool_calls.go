package transcript_cache

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"

func (c *Cache) ToolCalls(sessionIdentifier string) []tool_call.Call {
	modTime, size, exists := c.fileState(sessionIdentifier)

	if !exists {
		return nil
	}

	c.mutex.RLock()
	entry, found := c.calls[sessionIdentifier]
	c.mutex.RUnlock()

	if found && entry.ModTime.Equal(modTime) && entry.Size == size {
		return entry.Calls
	}

	calls := c.Client.ToolCalls(sessionIdentifier)
	c.mutex.Lock()
	c.calls[sessionIdentifier] = &callEntry{
		ModTime: modTime,
		Size:    size,
		Calls:   calls,
	}
	c.mutex.Unlock()

	return calls
}
