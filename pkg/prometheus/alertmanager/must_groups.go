package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) MustGroups() models.AlertGroups {
	result, e := c.Groups()
	errors.PanicOnError(e)

	return result
}
