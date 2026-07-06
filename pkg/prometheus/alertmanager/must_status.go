package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) MustStatus() *models.AlertmanagerStatus {
	result, e := c.Status()
	errors.PanicOnError(e)

	return result
}
