package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/maintenance"
)

func (c *Client) Maintenances() []maintenance.Maintenance {
	r, e := c.userClient.Maintenance.List(
		c.context,
		&maintenance.ListRequest{Type: maintenance.All},
	)
	errors.PanicOnError(e)

	return r.Maintenances
}
