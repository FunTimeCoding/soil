package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func profileLoads(
	sessionIdentifier string,
	c *tool_call.Call,
) []context_load.Load {
	tiers := profileEntries(resultText(c.Result))
	var result []context_load.Load

	for _, tier := range []string{constant.TierAlways, constant.TierRelevant} {
		for _, entry := range tiers[tier] {
			load := memoryLoad(sessionIdentifier, c, &entry)
			load.Tier = tier
			result = append(result, *load)
		}
	}

	return result
}
