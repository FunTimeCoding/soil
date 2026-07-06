package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/errors"
	rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AddResponderTeam(
	a *alert.Alert,
	team string,
) *rawAlert.AsyncAlertResult {
	result, e := c.teamClient.Alert.AddResponder(
		c.context,
		&rawAlert.AddResponderRequest{
			IdentifierType:  rawAlert.ALERTID,
			IdentifierValue: a.Identifier,
			Responder: rawAlert.Responder{
				Type: rawAlert.TeamResponder,
				Name: team,
			},
		},
	)
	errors.PanicOnError(e)

	return result
}
