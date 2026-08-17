package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"github.com/funtimecoding/soil/pkg/opnsense/state"
)

func (c *Client) States(phrase string) ([]*state.State, error) {
	rows, e := searchRows[response.State](c, constant.StateQuery, phrase)

	if e != nil {
		return nil, e
	}

	return state.NewSlice(rows), nil
}
