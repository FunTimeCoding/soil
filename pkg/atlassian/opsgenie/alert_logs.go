package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AlertLogs(identifier string) []alert.AlertLog {
	r, e := c.userClient.Alert.ListAlertLogs(
		c.context,
		&alert.ListAlertLogsRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return r.AlertLog
}
