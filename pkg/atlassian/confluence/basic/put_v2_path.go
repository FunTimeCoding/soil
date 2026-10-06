package basic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) PutV2Path(
	path string,
	body string,
	out any,
) error {
	return c.requester.Notation(
		request.New(http.MethodPut, path).WithBody(
			constant.Object,
			[]byte(body),
		),
		out,
	)
}
