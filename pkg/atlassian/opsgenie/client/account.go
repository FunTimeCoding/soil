package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/account"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

func Account(c *client.Config) *account.Client {
	result, e := account.NewClient(c)
	errors.PanicOnError(e)

	return result
}
