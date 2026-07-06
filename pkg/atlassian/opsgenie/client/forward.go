package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/forwarding_rule"
)

func Forward(c *client.Config) *forwarding_rule.Client {
	result, e := forwarding_rule.NewClient(c)
	errors.PanicOnError(e)

	return result
}
