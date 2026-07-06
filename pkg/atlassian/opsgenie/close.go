package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Close(identifier string) *alert.AsyncAlertResult {
	result, e := c.teamClient.Alert.Close(
		c.context,
		&alert.CloseAlertRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return result
}
