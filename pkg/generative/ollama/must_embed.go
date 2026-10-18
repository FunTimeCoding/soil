package ollama

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustEmbed(
	x context.Context,
	v []string,
) [][]float32 {
	result, e := c.Embed(x, v)
	errors.PanicOnError(e)

	return result
}
