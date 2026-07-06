package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/integration"
)

func Integration(c *client.Config) *integration.Client {
	result, e := integration.NewClient(c)
	errors.PanicOnError(e)

	return result
}
