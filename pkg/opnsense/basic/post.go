package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) Post(
	path string,
	body any,
	out any,
) error {
	return c.requester.Notation(
		request.New(http.MethodPost, path).WithNotation(body),
		out,
	)
}
