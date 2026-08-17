package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/forward"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) Forwards(phrase string) ([]*forward.Forward, error) {
	rows, e := searchRows[response.Forward](c, constant.ForwardSearch, phrase)

	if e != nil {
		return nil, e
	}

	return forward.NewSlice(rows), nil
}
