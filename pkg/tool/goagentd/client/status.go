package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
)

func (c *Client) Status() (*client.StatusResponse, error) {
	result, e := c.client.GetStatusWithResponse(c.context)

	if e != nil && Unreachable(e) {
		return nil, unreachable.Format("agent unreachable")
	}

	errors.PanicOnError(e)

	return result.JSON200, nil
}
