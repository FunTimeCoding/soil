package band

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustConsent() *response.Consent {
	result, e := c.Consent()
	errors.PanicOnError(e)

	return result
}
