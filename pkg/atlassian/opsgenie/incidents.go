package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
)

func (c *Client) Incidents(query string) []incident.Incident {
	r, e := c.userClient.Incident.List(
		c.context,
		&incident.ListRequest{Query: query},
	)
	errors.PanicOnError(e)

	return r.Incidents
}
