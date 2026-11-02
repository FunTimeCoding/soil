package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateSnippet(
	_ context.Context,
	_ *client.CreateSnippetParams,
	b client.CreateSnippetJSONRequestBody,
) (*client.SnippetCreateResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	c.Snippets[b.Name] = b.Content

	return &client.SnippetCreateResult{
		Status: "created",
		Volume: join.Empty("local:snippets/", b.Name),
	}, nil
}
