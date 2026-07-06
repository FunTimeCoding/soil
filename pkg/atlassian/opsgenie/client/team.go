package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
)

func Team(c *client.Config) *team.Client {
	result, e := team.NewClient(c)
	errors.PanicOnError(e)

	return result
}
