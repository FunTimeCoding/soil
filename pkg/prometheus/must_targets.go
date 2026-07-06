package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) MustTargets() v1.TargetsResult {
	result, e := c.Targets()
	errors.PanicOnError(e)

	return result
}
