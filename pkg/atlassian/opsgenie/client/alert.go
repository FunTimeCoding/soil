package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

func Alert(c *client.Config) *alert.Client {
	result, e := alert.NewClient(c)
	errors.PanicOnError(e)

	return result
}
