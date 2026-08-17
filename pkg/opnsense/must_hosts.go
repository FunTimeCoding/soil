package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/host"
)

func (c *Client) MustHosts(phrase string) []*host.Host {
	result, e := c.Hosts(phrase)
	errors.PanicOnError(e)

	return result
}
