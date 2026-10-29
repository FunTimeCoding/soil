package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostChunk(
	x context.Context,
	b client.PostChunkJSONRequestBody,
) (*client.ChunkPreview, error) {
	r, e := c.client.PostChunkWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("chunk", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
