package client

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/alert"
	"net/http"
)

func (c *Client) AlertsByIdentifier(
	identifiers []uint,
) ([]*alert.Alert, error) {
	if len(identifiers) == 0 {
		return []*alert.Alert{}, nil
	}

	parameter := make([]int, 0, len(identifiers))

	for _, i := range identifiers {
		parameter = append(parameter, int(i))
	}

	response, e := c.client.GetAlertsByIdentifierWithResponse(
		c.context,
		&client.GetAlertsByIdentifierParams{Identifier: parameter},
	)

	if e != nil {
		return nil, e
	}

	if response.StatusCode() != http.StatusOK || response.JSON200 == nil {
		return nil, unexpected.Format(
			"alerts by identifier status %d",
			response.StatusCode(),
		)
	}

	return toAlerts(*response.JSON200)
}
