package opnsense

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustReconfigureDnsmasq() {
	errors.PanicOnError(c.ReconfigureDnsmasq())
}
