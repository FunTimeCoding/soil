package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetSnippet(
	_ context.Context,
	name string,
	_ *client.GetSnippetParams,
) (*client.SnippetDetail, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	content, okay := c.Snippets[name]

	if !okay {
		return nil, not_found.Format("snippet %s", name)
	}

	return &client.SnippetDetail{Name: name, Content: content}, nil
}
