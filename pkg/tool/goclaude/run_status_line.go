package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunStatusLine(
	c *client.ClientWithResponses,
	body []byte,
) string {
	input := parseStatusLineInput(body)
	reportContext(c, input)

	return fmt.Sprintf(
		"%d%% context",
		int(input.ContextWindow.UsedPercentage),
	)
}
