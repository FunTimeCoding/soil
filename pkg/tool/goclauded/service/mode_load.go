package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func modeLoad(
	sessionIdentifier string,
	c *tool_call.Call,
	tag string,
) *context_load.Load {
	result := context_load.New()
	result.SessionIdentifier = sessionIdentifier
	result.CallIdentifier = c.Identifier
	result.Reference = tag
	result.Kind = constant.LoadKindMode
	result.Name = tag
	result.OccurredAt = callTime(c)

	return result
}
