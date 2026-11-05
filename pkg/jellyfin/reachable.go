package jellyfin

import (
	"context"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
)

func (c *Client) Reachable(context.Context) error {
	var out map[string]any

	return c.basic.Get(constant.ReachablePath, nil, &out)
}
