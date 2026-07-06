package ollama

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) MustList() []api.ListModelResponse {
	result, e := c.List()
	errors.PanicOnError(e)

	return result
}
