package basic

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) classified(q *request.Request) *request.Request {
	if c.rejecting {
		q.WithStatus(http.StatusBadRequest, errorsConstant.ValidationClass)
	}

	return q
}
