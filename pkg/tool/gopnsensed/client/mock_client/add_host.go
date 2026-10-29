package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"strconv"
)

func (c *Client) AddHost(
	_ context.Context,
	h client.HostRequest,
	apply bool,
) (*client.Identifier, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	identifier := strconv.Itoa(len(c.HostList) + 1)
	c.HostList = append(c.HostList, host(identifier, h))
	c.reconfigureWhen(apply)

	return &client.Identifier{Identifier: identifier}, nil
}
