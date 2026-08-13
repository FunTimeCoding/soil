package ollama

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustEmbedSingle(v string) []float32 {
	result, e := c.EmbedSingle(v)
	errors.PanicOnError(e)

	return result
}
