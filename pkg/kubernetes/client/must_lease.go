package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
)

func (c *Client) MustLease(
	namespace string,
	name string,
) *lease.Lease {
	result, e := c.Lease(namespace, name)
	errors.PanicOnError(e)

	return result
}
