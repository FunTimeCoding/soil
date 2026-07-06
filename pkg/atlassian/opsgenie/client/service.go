package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/service"
)

func Service(c *client.Config) *service.Client {
	result, e := service.NewClient(c)
	errors.PanicOnError(e)

	return result
}
