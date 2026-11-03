package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func (c *Client) CreateView(
	_ context.Context,
	b client.CreateViewJSONRequestBody,
) (*client.View, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	result := client.View{ViewId: c.next, Title: b.Title}

	if b.Content != "" {
		result.Text = new(b.Content)
	}

	if b.Syntax != nil {
		result.Syntax = *b.Syntax
	}

	c.next++
	c.ViewList = append(c.ViewList, result)

	return &result, nil
}
