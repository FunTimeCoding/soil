package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
)

func (c *Client) MustInterfaces() []*network_interface.Interface {
	result, e := c.Interfaces()
	errors.PanicOnError(e)

	return result
}
