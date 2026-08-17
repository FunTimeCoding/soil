package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func memoryLoads(
	sessionIdentifier string,
	c *tool_call.Call,
) []context_load.Load {
	var result []context_load.Load

	for _, entry := range memoryEntries(resultText(c.Result)) {
		result = append(result, *memoryLoad(sessionIdentifier, c, &entry))
	}

	return result
}
