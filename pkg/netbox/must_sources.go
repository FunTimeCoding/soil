package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/source"
)

func (c *Client) MustSources() []*source.Source {
	result, e := c.Sources()
	errors.PanicOnError(e)

	return result
}
