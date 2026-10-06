package basic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) PostV2Path(
	path string,
	body string,
	out any,
) error {
	return c.requester.Notation(
		request.New(http.MethodPost, path).WithBody(
			constant.Object,
			[]byte(body),
		),
		out,
	)
}
