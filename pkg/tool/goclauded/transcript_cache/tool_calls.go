package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache/entry"
)

func (c *Cache) ToolCalls(sessionIdentifier string) []*tool_call.Call {
	modTime, size, exists := c.fileState(sessionIdentifier)

	if !exists {
		return nil
	}

	c.mutex.RLock()
	e, found := c.calls[sessionIdentifier]
	c.mutex.RUnlock()

	if found && e.ModTime.Equal(modTime) && e.Size == size {
		return e.Calls
	}

	calls := c.Client.ToolCalls(sessionIdentifier)
	c.mutex.Lock()
	c.calls[sessionIdentifier] = entry.NewCall(modTime, size, calls)
	c.mutex.Unlock()

	return calls
}
