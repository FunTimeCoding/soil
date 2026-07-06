package ollama

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustEmbed(
	model string,
	v []string,
) [][]float32 {
	result, e := c.Embed(model, v)
	errors.PanicOnError(e)

	return result
}
