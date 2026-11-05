package basic

import (
	"encoding/json"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/habitica/envelope"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (c *Client) unwrap(
	q *request.Request,
	out any,
) error {
	if c.rejecting {
		q.WithStatus(http.StatusBadRequest, errorsConstant.ValidationClass)
	}

	var v envelope.Envelope

	if e := c.requester.Notation(q, &v); e != nil {
		return e
	}

	return json.Unmarshal(v.Payload, out)
}
