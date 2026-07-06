package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func RunRegister(
	c *client.ClientWithResponses,
	session string,
) string {
	response, e := c.PostRegisterWithResponse(
		context.Background(),
		client.PostRegisterJSONRequestBody{Session: session},
	)

	if e != nil {
		return ""
	}

	if response.JSON200 == nil {
		return ""
	}

	return response.JSON200.Callsign
}
