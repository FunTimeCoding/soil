package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
)

func (c *Client) MustSilenceByRule(name string) *silence.Silence {
	result, e := c.SilenceByRule(name)
	errors.PanicOnError(e)

	return result
}
