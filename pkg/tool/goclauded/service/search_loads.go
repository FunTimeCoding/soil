package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func searchLoads(
	sessionIdentifier string,
	c *tool_call.Call,
) []context_load.Load {
	query := inputString(c.Input, generative.ParameterQuery)
	var result []context_load.Load

	for _, entry := range memoryEntries(resultText(c.Result)) {
		load := memoryLoad(sessionIdentifier, c, &entry)
		load.Kind = constant.LoadKindSearch
		load.Query = query
		result = append(result, *load)
	}

	return result
}
