package aptly

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Reachable(_ context.Context) error {
	var result map[string]string

	return c.requester.Notation(request.Get("/api/version"), &result)
}
