package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (c *Client) PostDocument(
	_ context.Context,
	b client.PostDocumentJSONRequestBody,
) (*client.PushOutcome, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	c.Pushed = append(c.Pushed, b)
	body := b.Body
	c.documents[key(b.Collection, b.Path)] = &client.SearchResult{
		Collection:  b.Collection,
		Path:        b.Path,
		Body:        &body,
		Hash:        store.HashContent(b.Body),
		Metadata:    b.Metadata,
		FilePath:    key(b.Collection, b.Path),
		VirtualPath: key(b.Collection, b.Path),
	}

	return &client.PushOutcome{}, nil
}
