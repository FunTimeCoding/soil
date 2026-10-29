package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) SetHost(
	_ context.Context,
	identifier string,
	h client.HostRequest,
	apply bool,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	i := c.index(identifier)

	if i < 0 {
		return not_found.New(constant.HostSubject, identifier)
	}

	c.HostList[i] = host(identifier, h)
	c.reconfigureWhen(apply)

	return nil
}
