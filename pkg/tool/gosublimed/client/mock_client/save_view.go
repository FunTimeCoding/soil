package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) SaveView(
	_ context.Context,
	identifier int,
	b client.SaveViewJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return e
	}

	i, e := c.index(identifier)

	if e != nil {
		return e
	}

	c.ViewList[i].FilePath = ""

	if b.FilePath != nil {
		c.ViewList[i].FilePath = *b.FilePath
	}

	c.ViewList[i].IsDirty = false

	return nil
}
