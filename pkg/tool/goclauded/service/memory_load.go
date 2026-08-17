package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"strconv"
)

func memoryLoad(
	sessionIdentifier string,
	c *tool_call.Call,
	entry *memoryEntry,
) *context_load.Load {
	result := context_load.New()
	result.SessionIdentifier = sessionIdentifier
	result.CallIdentifier = c.Identifier
	result.Reference = strconv.FormatInt(entry.Identifier, 10)
	result.Kind = constant.LoadKindMemory
	result.Name = entry.Name
	result.OccurredAt = callTime(c)

	return result
}
