package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) Reconcile(
	full bool,
	write bool,
) *response.Response {
	result, e := c.client.ReconcileWithResponse(
		c.context,
		&client.ReconcileParams{
			Instance: &c.instance,
			Full:     &full,
			Write:    &write,
		},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
