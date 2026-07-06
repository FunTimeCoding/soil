package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
)

func Incident(c *client.Config) *incident.Client {
	result, e := incident.NewClient(c)
	errors.PanicOnError(e)

	return result
}
