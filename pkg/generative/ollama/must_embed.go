package ollama

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustEmbed(v []string) [][]float32 {
	result, e := c.Embed(v)
	errors.PanicOnError(e)

	return result
}
