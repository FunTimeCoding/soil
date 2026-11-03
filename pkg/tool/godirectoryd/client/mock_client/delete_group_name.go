package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"slices"
)

func (c *Client) DeleteGroupName(
	_ context.Context,
	name string,
) (*client.Receipt, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	i, e := c.groupIndex(name)

	if e != nil {
		return nil, e
	}

	c.GroupList = slices.Delete(c.GroupList, i, i+1)

	return &client.Receipt{Message: fmt.Sprintf("deleted group: %s", name)}, nil
}
