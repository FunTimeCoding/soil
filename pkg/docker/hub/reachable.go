package hub

import (
	"context"
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Reachable(_ context.Context) error {
	var result map[string]any

	return c.requester.Notation(request.Get(constant.ProbeImage), &result)
}
