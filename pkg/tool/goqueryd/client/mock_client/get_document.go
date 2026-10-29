package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) GetDocument(
	_ context.Context,
	p *client.GetDocumentParams,
) (*client.SearchResult, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	d, okay := c.documents[p.Path]

	if !okay {
		return nil, not_found.New(constant.DocumentKind, p.Path)
	}

	result := *d

	return &result, nil
}
