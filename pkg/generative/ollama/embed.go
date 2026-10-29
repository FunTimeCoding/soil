package ollama

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/ollama/refusal"
	"github.com/ollama/ollama/api"
)

func (c *Client) Embed(
	x context.Context,
	v []string,
) ([][]float32, error) {
	result, e := c.client.Embed(x, &api.EmbedRequest{Model: c.model, Input: v})

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return result.Embeddings, nil
}
