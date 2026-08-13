package openai

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/openai/openai-go/v3"
)

func (c *Client) Embed(v []string) ([][]float32, error) {
	response, e := c.client.Embeddings.New(
		context.Background(),
		openai.EmbeddingNewParams{
			Model: constant.OpenAIEmbedModel,
			Input: openai.EmbeddingNewParamsInputUnion{OfArrayOfStrings: v},
		},
	)

	if e != nil {
		return nil, e
	}

	result := make([][]float32, len(response.Data))

	for _, d := range response.Data {
		vector := make([]float32, len(d.Embedding))

		for i, value := range d.Embedding {
			vector[i] = float32(value)
		}

		result[d.Index] = vector
	}

	return result, nil
}
