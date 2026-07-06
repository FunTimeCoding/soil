package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/escalation"
)

func Escalation(c *client.Config) *escalation.Client {
	result, e := escalation.NewClient(c)
	errors.PanicOnError(e)

	return result
}
