package transcript

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func profileLoads(
	sessionIdentifier string,
	c *tool_call.Call,
) ([]context_load.Load, bool) {
	tiers, readable := profileEntries(resultText(c.Result))

	return tierLoads(sessionIdentifier, c, tiers), readable
}
