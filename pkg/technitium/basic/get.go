package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/envelope"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Get(path string) (json.RawMessage, error) {
	var v envelope.Envelope

	if e := c.requester.Notation(
		request.Absolute(join.Empty(c.base, path)),
		&v,
	); e != nil {
		return nil, e
	}

	if v.Status != constant.OkayStatus {
		return nil, unexpected.Format("technitium %s: %s", v.Status, v.Message)
	}

	return v.Payload, nil
}
