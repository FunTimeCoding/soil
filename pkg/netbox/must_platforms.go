package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
)

func (c *Client) MustPlatforms() []*platform.Platform {
	result, e := c.Platforms()
	errors.PanicOnError(e)

	return result
}
