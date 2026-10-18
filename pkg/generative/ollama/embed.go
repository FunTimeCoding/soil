package ollama

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/ollama/ollama/api"
)

func (c *Client) Embed(
	x context.Context,
	v []string,
) ([][]float32, error) {
	result, e := c.client.Embed(
		x,
		&api.EmbedRequest{Model: constant.OllamaEmbedModel, Input: v},
	)

	if e != nil {
		return nil, e
	}

	return result.Embeddings, nil
}
