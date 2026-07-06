package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/errors"
	rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Alert(
	identifier string,
	t *team.Team,
) *alert.Alert {
	result, e := c.userClient.Alert.Get(
		c.context,
		&rawAlert.GetAlertRequest{
			IdentifierType:  rawAlert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return c.processor().ProcessOne(
		alert.NewDetail(result, c.AlertOption(), t),
	)
}
