package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"slices"
)

func (c *Client) DeleteHost(
	_ context.Context,
	identifier string,
	apply bool,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	i := c.index(identifier)

	if i < 0 {
		return not_found.New(constant.HostSubject, identifier)
	}

	c.HostList = slices.Delete(c.HostList, i, i+1)
	c.reconfigureWhen(apply)

	return nil
}
