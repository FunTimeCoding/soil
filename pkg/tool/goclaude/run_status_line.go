package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunStatusLine(
	c *client.ClientWithResponses,
	body []byte,
) string {
	input := parseStatusLineInput(body)
	reportContext(c, input)
	context := fmt.Sprintf("%d%%", int(input.ContextWindow.UsedPercentage))
	model := shortModelName(input.Model.DisplayName)

	if model == "" {
		return context
	}

	return join.Space(model, context)
}
