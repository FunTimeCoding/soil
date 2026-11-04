package basic

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) Post(
	x context.Context,
	path string,
	body any,
	out any,
) error {
	return c.requester.Notation(
		request.New(http.MethodPost, path).WithContext(x).WithNotation(body),
		out,
	)
}
