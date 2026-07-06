package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/deployment"
)

func Deployment(c *client.Config) *deployment.Client {
	result, e := deployment.NewClient(c)
	errors.PanicOnError(e)

	return result
}
