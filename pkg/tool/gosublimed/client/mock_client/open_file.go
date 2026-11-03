package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) OpenFile(
	_ context.Context,
	b client.OpenFileJSONRequestBody,
) (*client.View, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	result := client.View{ViewId: c.next, FilePath: b.FilePath}
	c.next++
	c.ViewList = append(c.ViewList, result)

	return &result, nil
}
