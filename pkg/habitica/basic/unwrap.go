package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/habitica/envelope"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) unwrap(
	q *request.Request,
	out any,
) error {
	var v envelope.Envelope

	if e := c.requester.Notation(q, &v); e != nil {
		return e
	}

	return json.Unmarshal(v.Payload, out)
}
