package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Create(
	email string,
	text string,
) *alert.AsyncAlertResult {
	result, e := c.teamClient.Alert.Create(
		c.context,
		&alert.CreateAlertRequest{User: email, Message: text},
	)
	errors.PanicOnError(e)

	return result
}
