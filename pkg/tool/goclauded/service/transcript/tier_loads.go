package transcript

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/memory_payload"
)

func tierLoads(
	sessionIdentifier string,
	c *tool_call.Call,
	tiers map[string][]memory_payload.Entry,
) []context_load.Load {
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
