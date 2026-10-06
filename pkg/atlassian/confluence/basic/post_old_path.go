package basic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) PostOldPath(
	path string,
	body string,
) error {
	_, e := c.old.Bytes(
		request.New(http.MethodPost, path).WithBody(
			constant.Object,
			[]byte(body),
		),
	)

	return e
}
