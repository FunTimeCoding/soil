package goclaude

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"time"
)

func channelCallsign(
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

	return AwaitCallsign(c.LongClient(), session, since, interval)
}
