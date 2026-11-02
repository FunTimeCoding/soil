package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeleteSnippet(
	_ context.Context,
	name string,
	_ *client.DeleteSnippetParams,
) (*client.StatusResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if _, okay := c.Snippets[name]; !okay {
		return nil, not_found.Format("snippet %s", name)
	}

	delete(c.Snippets, name)

	return &client.StatusResult{Status: "deleted"}, nil
}
