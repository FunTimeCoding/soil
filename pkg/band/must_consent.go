package band

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustConsent() *ConsentResponse {
	result, e := c.Consent()
	errors.PanicOnError(e)

	return result
}
