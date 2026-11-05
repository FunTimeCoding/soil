package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func (c *Client) GetLog(
	_ context.Context,
	name string,
	all bool,
) (*client.Log, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if _, e := c.index(name); e != nil {
		return nil, e
	}

	result := &client.Log{Lines: c.Logs[name]}

	if !all {
		result.Older = c.Older[name]
	}

	return result, nil
}
