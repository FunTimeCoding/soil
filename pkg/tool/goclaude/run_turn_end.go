package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunTurnEnd(
	c *client.ClientWithResponses,
	session string,
) {
	if _, e := c.PostTurnEndWithResponse(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: session},
	); e != nil {
		return
	}
}
