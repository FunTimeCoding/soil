package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/contact"
)

func Contact(c *client.Config) *contact.Client {
	result, e := contact.NewClient(c)
	errors.PanicOnError(e)

	return result
}
