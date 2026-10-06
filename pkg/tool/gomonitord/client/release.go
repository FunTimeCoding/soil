package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"net/http"
)

func (c *Client) Release(
	item string,
	owner string,
) error {
	r, e := c.client.DeleteClaimWithResponse(
		c.context,
		&client.DeleteClaimParams{Item: item, Owner: owner},
	)

	if e != nil {
		return fmt.Errorf("release %s: %w", item, e)
	}

	if r.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("release %s: %s", item, r.Status())
	}

	return nil
}
