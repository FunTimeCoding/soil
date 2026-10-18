package ollama

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustEmbedSingle(
	x context.Context,
	v string,
) []float32 {
	result, e := c.EmbedSingle(x, v)
	errors.PanicOnError(e)

	return result
}
