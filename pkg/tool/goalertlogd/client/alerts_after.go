package client

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/alert"
	"net/http"
)

func (c *Client) AlertsAfter(
	identifier uint,
	limit int,
) ([]*alert.Alert, error) {
	response, e := c.client.GetAlertsAfterWithResponse(
		c.context,
		&client.GetAlertsAfterParams{
			Identifier: int(identifier),
			Limit:      &limit,
		},
	)

	if e != nil {
		return nil, e
	}

	if response.StatusCode() != http.StatusOK || response.JSON200 == nil {
		return nil, unexpected.Format(
			"alerts after status %d",
			response.StatusCode(),
		)
	}

	return toAlerts(*response.JSON200)
}
