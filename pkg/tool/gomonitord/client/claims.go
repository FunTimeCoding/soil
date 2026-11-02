package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"net/http"
)

func (c *Client) Claims() ([]client.Claim, error) {
	r, e := c.client.GetClaimsWithResponse(c.context)

	if e != nil {
		return nil, fmt.Errorf("claims: %w", e)
	}

	if r.StatusCode() != http.StatusOK || r.JSON200 == nil {
		return nil, fmt.Errorf("claims: %s", r.Status())
	}

	return *r.JSON200, nil
}
