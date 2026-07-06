package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) MustStatus() v1.RuntimeinfoResult {
	result, e := c.Status()
	errors.PanicOnError(e)

	return result
}
