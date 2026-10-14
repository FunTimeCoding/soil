package client

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
)

func (c *Client) LatestIdentifier() (uint, error) {
	response, e := c.client.GetStatusWithResponse(c.context)

	if e != nil {
		return 0, e
	}

	if response.StatusCode() != http.StatusOK || response.JSON200 == nil {
		return 0, unexpected.Format("status %d", response.StatusCode())
	}

	return uint(response.JSON200.LatestIdentifier), nil
}
