package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/policy"
)

func Policy(c *client.Config) *policy.Client {
	result, e := policy.NewClient(c)
	errors.PanicOnError(e)

	return result
}
