package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
)

func (c *Client) MustSilences(expired bool) []*silence.Silence {
	result, e := c.Silences(expired)
	errors.PanicOnError(e)

	return result
}
