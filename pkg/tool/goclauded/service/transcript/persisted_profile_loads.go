package transcript

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"strings"
)

func persistedProfileLoads(
	sessionIdentifier string,
	c *tool_call.Call,
) []context_load.Load {
	path := inputString(c.Input, constant.FilePath)

	if !strings.Contains(path, constant.ToolResultDirectory) {
		return nil
	}

	saved := numberedText(resultText(c.Result))
	tiers, _ := profileEntries(resultText(saved))

	return tierLoads(sessionIdentifier, c, tiers)
}
