package utilization

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Reachable(x context.Context) error {
	response, e := c.requester.Send(
		request.Get(webConstant.RootPath).WithContext(x),
	)

	if e == nil {
		errors.LogClose(response.Body)

		return nil
	}

	if unreachable.Is(e) || timeout.Is(e) || dropped.Is(e) {
		return e
	}

	return nil
}
