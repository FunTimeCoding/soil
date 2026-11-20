package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) GetSystemFindings(
	_ context.Context,
	name string,
) (*client.SystemDetail, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	result := c.detail(name)

	if result == nil {
		return nil, not_found.New(constant.SystemKind, name)
	}

	return result, nil
}
