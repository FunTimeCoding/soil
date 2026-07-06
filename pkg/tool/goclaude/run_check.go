package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunCheck(
	c *client.ClientWithResponses,
	session string,
) string {
	response, e := c.GetCheckWithResponse(
		context.Background(),
		&client.GetCheckParams{Session: session},
	)

	if e != nil {
		return ""
	}

	if response.JSON200 == nil {
		return ""
	}

	if !response.JSON200.Changed {
		return ""
	}

	return formatCheckContext(response.JSON200)
}
