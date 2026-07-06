package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
)

func Heartbeat(c *client.Config) *heartbeat.Client {
	result, e := heartbeat.NewClient(c)
	errors.PanicOnError(e)

	return result
}
