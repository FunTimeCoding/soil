package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) MustReceivers() []*models.Receiver {
	result, e := c.Receivers()
	errors.PanicOnError(e)

	return result
}
