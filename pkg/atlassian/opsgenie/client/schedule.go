package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func Schedule(c *client.Config) *schedule.Client {
	result, e := schedule.NewClient(c)
	errors.PanicOnError(e)

	return result
}
