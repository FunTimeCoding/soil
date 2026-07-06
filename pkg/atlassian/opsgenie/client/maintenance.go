package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/maintenance"
)

func Maintenance(c *client.Config) *maintenance.Client {
	result, e := maintenance.NewClient(c)
	errors.PanicOnError(e)

	return result
}
