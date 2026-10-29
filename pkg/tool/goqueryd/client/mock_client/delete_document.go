package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) DeleteDocument(
	_ context.Context,
	p *client.DeleteDocumentParams,
) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return false, c.Failure
	}

	c.Deleted = append(c.Deleted, *p)
	k := key(p.Collection, p.Path)
	_, okay := c.documents[k]
	delete(c.documents, k)

	return okay, nil
}
