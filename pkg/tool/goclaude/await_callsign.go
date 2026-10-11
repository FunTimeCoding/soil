package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
	"time"
)

func AwaitCallsign(
	c *client.ClientWithResponses,
	session string,
	since time.Time,
	interval time.Duration,
) string {
	for failures := 0; failures < constant.ChannelCallsignAttempts; {
		response, e := c.GetChannelCallsignWithResponse(
			context.Background(),
			&client.GetChannelCallsignParams{Session: session, Since: since},
		)

		if e == nil && response.JSON200 != nil {
			return response.JSON200.Callsign
		}

		if e == nil && response.StatusCode() == http.StatusNoContent {
			continue
		}

		failures++
		time.Sleep(interval)
	}

	return ""
}
