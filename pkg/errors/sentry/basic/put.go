package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) Put(
	path string,
	body any,
	out any,
) error {
	return c.requester.Notation(
		request.New(http.MethodPut, path).WithNotation(body),
		out,
	)
}
