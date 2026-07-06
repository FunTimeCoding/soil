package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/errors"
	rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AddResponderUser(
	a *alert.Alert,
	user string,
) {
	_, e := c.teamClient.Alert.AddResponder(
		c.context,
		&rawAlert.AddResponderRequest{
			IdentifierType:  rawAlert.ALERTID,
			IdentifierValue: a.Identifier,
			Responder: rawAlert.Responder{
				Type: rawAlert.UserResponder,
				Name: user,
			},
		},
	)
	errors.PanicOnError(e)
}
