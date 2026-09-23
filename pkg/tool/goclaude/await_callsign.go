package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"time"
)

func awaitCallsign(
	c *command_context.Context,
	since time.Time,
	interval time.Duration,
) string {
	session := environment.Optional(
		constant.HarnessSessionIdentifierEnvironment,
	)

	if session == "" {
		return ""
	}

	for failures := 0; failures < constant.ChannelCallsignAttempts; {
		response, e := c.Client().GetChannelCallsignWithResponse(
			context.Background(),
			&client.GetChannelCallsignParams{Session: session, Since: since},
		)

		if e == nil && response.JSON200 != nil {
			return response.JSON200.Callsign
		}

		failures++
		time.Sleep(interval)
	}

	return ""
}
