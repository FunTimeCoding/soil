package opsgenie

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/client"
	"github.com/funtimecoding/soil/pkg/errors"
)

func New(
	userKey string,
	teamKey string,
	webHost string,
) *Client {
	errors.FatalOnEmpty(userKey, "userKey")
	errors.FatalOnEmpty(webHost, "webHost")
	result := &Client{
		context:    context.Background(),
		webHost:    webHost,
		userClient: client.New(userKey),
	}

	if teamKey != "" {
		result.teamClient = client.New(teamKey)
	}

	return result
}
