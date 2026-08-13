package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/ollama/ollama/api"
)

func (c *Client) Embed(v []string) ([][]float32, error) {
	result, e := c.client.Embed(
		c.context,
		&api.EmbedRequest{Model: constant.OllamaEmbedModel, Input: v},
	)

	if e != nil {
		return nil, e
	}

	return result.Embeddings, nil
}
