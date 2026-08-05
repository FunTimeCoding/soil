package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunStatusline(
	c *client.ClientWithResponses,
	body []byte,
) string {
	input := parseStatuslineInput(body)
	reportContext(c, input)

	return fmt.Sprintf(
		"%d%% context",
		int(input.ContextWindow.UsedPercentage),
	)
}
