package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunSessionEnd(
	c *client.ClientWithResponses,
	session string,
	reason string,
) {
	body := client.PostSessionEndJSONRequestBody{Session: session}

	if reason != "" {
		body.Reason = &reason
	}

	if _, e := c.PostSessionEndWithResponse(
		context.Background(),
		body,
	); e != nil {
		return
	}
}
