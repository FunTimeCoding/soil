package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
)

func (c *Client) DeleteSystem(
	_ context.Context,
	name string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return c.failure
	}

	if c.detail(name) == nil {
		return not_found.New(constant.SystemKind, name)
	}

	delete(c.snapshots, name)
	delete(c.updated, name)
	delete(c.details, name)

	return nil
}
