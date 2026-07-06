package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/notification"
)

func Notification(c *client.Config) *notification.Client {
	result, e := notification.NewClient(c)
	errors.PanicOnError(e)

	return result
}
