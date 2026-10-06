package basic

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
	"net/url"
)

func (c *Client) Post(
	path string,
	v url.Values,
	body any,
	out any,
) error {
	q := request.New(http.MethodPost, path).WithParameters(v)

	if body != nil {
		q.WithNotation(body)
	}

	return c.unwrap(q, out)
}
