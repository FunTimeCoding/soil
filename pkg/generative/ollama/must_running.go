package ollama

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) MustRunning() []api.ProcessModelResponse {
	result, e := c.Running()
	errors.PanicOnError(e)

	return result
}
