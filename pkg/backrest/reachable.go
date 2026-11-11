package backrest

import (
	"context"
	"github.com/funtimecoding/soil/pkg/backrest/constant"
)

func (c *Client) Reachable(x context.Context) error {
	_, e := c.run(x, constant.VersionCommand)

	return e
}
