package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/integration"
)

func (c *Client) Integrations() []integration.GenericFields {
	r, e := c.userClient.Integration.List(c.context)
	errors.PanicOnError(e)

	return r.Integrations
}
