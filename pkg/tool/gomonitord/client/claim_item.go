package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"net/http"
)

func (c *Client) ClaimItem(
	item string,
	owner string,
) error {
	r, e := c.client.PutClaimWithResponse(
		c.context,
		client.PutClaimJSONRequestBody{Item: item, Owner: owner},
	)

	if e != nil {
		return fmt.Errorf("claim %s: %w", item, e)
	}

	if r.StatusCode() != http.StatusOK {
		return fmt.Errorf("claim %s: %s", item, r.Status())
	}

	return nil
}
