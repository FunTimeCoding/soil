package basic

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) Post(
	path string,
	body []byte,
) error {
	if c.verbose {
		console.Format("POST %s\n%s\n", path, body)
	}

	_, e := c.requester.Bytes(
		request.New(http.MethodPost, path).WithBody(constant.Object, body),
	)

	return e
}
